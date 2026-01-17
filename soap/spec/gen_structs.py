import sys
import re
from xml.etree import ElementTree as ET
from collections import OrderedDict

def camel_case(s):
    parts = re.split('[_-]', s)
    return ''.join(p.capitalize() for p in parts if p)

def replace_ns(name, old_ns, new_ns):
    # Replace the namespace prefix in tag or attribute
    if ':' in name:
        return name.replace(old_ns + ':', new_ns + ':')
    return name

def parse_type(tag, comment):
    is_array = 'Zero or more repetitions' in comment or '1 or more repetitions' in comment
    is_pointer = True  # All fields are pointers
    return is_array, is_pointer

def parse_element(elem, parent_comments=None, structs=None, struct_name=None, old_ns=None, new_ns=None):
    if structs is None:
        structs = OrderedDict()
    fields = []
    # Add attributes for this element
    for attr in elem.attrib:
        attr_name = attr.split(':')[-1] if ':' in attr else attr
        go_field = camel_case(attr_name)
        ns_attr = replace_ns(attr, old_ns, new_ns)
        fields.append((go_field, '*string', False, True, None, f'xml:"{ns_attr},attr"', None))
    # Add child elements
    for child in elem:
        comment = ''
        if child.tail and child.tail.strip().startswith('<!--'):
            comment = child.tail.strip()
        elif child.text and child.text.strip().startswith('<!--'):
            comment = child.text.strip()
        elif parent_comments:
            comment = parent_comments.pop(0)
        else:
            comment = ''
        comment = re.sub(r'<!--|-->', '', comment).strip()
        tag = re.sub(r'^{.*}', '', child.tag)
        tag_ns, tag_name = (tag.split(':', 1) if ':' in tag else (None, tag))
        tag = replace_ns(tag, old_ns, new_ns)
        is_array, is_pointer = parse_type(tag_name, comment)
        go_type = 'string'
        child_attrs = list(child.attrib.keys())
        if list(child) or child_attrs:
            nested_struct_name = camel_case(tag_name)
            go_type = nested_struct_name
            if nested_struct_name not in structs:
                nested_fields, _ = parse_element(child, structs=structs, struct_name=nested_struct_name, old_ns=old_ns, new_ns=new_ns)
                structs[nested_struct_name] = nested_fields
            fields.append((tag_name, f'*{go_type}' if not is_array else f'*[]{go_type}', is_array, is_pointer, comment, f'xml:"{tag},omitempty"', child.attrib))
        else:
            fields.append((tag_name, '*string', is_array, is_pointer, comment, f'xml:"{tag},omitempty"', child.attrib))
    return fields, None

def gen_struct(name, fields, namespace, is_root=False, root_tag=None):
    lines = []
    lines.append(f'type {name} struct ' + '{')
    if is_root:
        lines.append(f'    XMLName xml.Name `xml:"{namespace}:{root_tag}"`')
        lines.append(f'    XMLNamespace *string `xml:"xmlns,attr"`')
    used_fields = set()
    for field in fields:
        tag, go_type, is_array, is_pointer, comment, xml_tag, attrs = field
        field_name = camel_case(tag)
        if field_name in used_fields:
            continue
        used_fields.add(field_name)
        comment_str = f'// {comment}' if comment else ''
        lines.append(f'    {field_name} {go_type} `{xml_tag}` {comment_str}')
        # Add attribute fields for child element if any
        if attrs:
            for attr in attrs:
                attr_name = attr.split(':')[-1] if ':' in attr else attr
                attr_field = camel_case(attr_name)
                if attr_field in used_fields:
                    continue
                used_fields.add(attr_field)
                ns_attr = replace_ns(attr, namespace, namespace)
                lines.append(f'    {attr_field} *string `xml:"{ns_attr},attr"`')
    lines.append('}')
    return lines

def extract_comments(xml_lines):
    comments = []
    for line in xml_lines:
        if '<!--' in line:
            comment = re.search(r'<!--(.*?)-->', line)
            if comment:
                comments.append(comment.group(1).strip())
    return comments

def main(xml_path, namespace, package):
    with open(xml_path) as f:
        xml_lines = f.readlines()
    comments = extract_comments(xml_lines)
    tree = ET.parse(xml_path)
    root = tree.getroot()
    root_tag = re.sub(r'^{.*}', '', root.tag)
    root_ns, root_name = (root_tag.split(':', 1) if ':' in root_tag else (None, root_tag))
    structs = OrderedDict()
    fields, _ = parse_element(root, comments, structs, struct_name=camel_case(root_name), old_ns=root_ns, new_ns=namespace)
    structs[camel_case(root_name)] = fields

    print(f'package {package}\n')
    print('import "encoding/xml"\n')
    for i, (name, fields) in enumerate(structs.items()):
        is_root = (i == len(structs) - 1)
        print('\n'.join(gen_struct(
            name, fields, namespace,
            is_root=is_root,
            root_tag=root_name if is_root else None
        )))
        print()

if __name__ == '__main__':
    if len(sys.argv) != 4:
        print("Usage: python generate_go_structs.py <xml_file> <namespace> <package>")
        sys.exit(1)
    main(sys.argv[1], sys.argv[2], sys.argv[3])