import sys
import glob
from xml.etree import ElementTree as ET
from collections import OrderedDict

def go_name(s: str):
    no_namespace = s.split('}')[-1] if '}' in s else s
    no_namespace = no_namespace.replace('-', '_')
    parts = no_namespace.split('_')
    go_field = ''.join(part.capitalize() for part in parts)
    # Check if the go_field starts with an unallowed golang variable character
    if go_field and not go_field[0].isalpha():
        go_field = '_' + go_field
    return go_field, no_namespace

def parse_element(elem: ET.Element, namespace: str, structs: dict, is_root: bool=False):
    go_field, xml_name = go_name(elem.tag)

    if structs.get(go_field):
        fields = structs[go_field]
    else:
        fields = dict[str, list[tuple[str, str]]]()
        structs[go_field] = fields

    if elem.text is not None and elem.text.strip():
        fields['Value'] = ('*string', 'xml:",chardata"')

    if is_root:
        fields['XMLName'] = ('xml.Name', f'xml:"{namespace}:{xml_name}"')
        fields['XMLNamespace'] = ('*string', f'xml:"xmlns:{namespace},attr,omitempty"')

    # Add attributes for this element
    for _, (name, _) in enumerate(elem.attrib.items()):
        go_field, xml_name = go_name(name)
        fields[go_field] = ('*string', f'xml:"{namespace}:{xml_name},attr,omitempty"')

    # Add child fields
    for child in elem:
        go_field, xml_name = go_name(child.tag)
        fields[go_field] = (f'*{go_field}', f'xml:"{namespace}:{xml_name},omitempty"')
        parse_element(child, namespace, structs, False)

def gen_struct(name, fields):
    lines = []
    lines.append(f'type {name} struct ' + '{')
    for go_field, (go_type, xml_tag) in fields.items():
        lines.append(f'    {go_field} {go_type} `{xml_tag}`')
    lines.append('}')
    return lines

def main(xml_path, package):
    files = glob.glob(xml_path)
    structs = OrderedDict()
    for file in files:
        tree = ET.parse(file)
        parse_element(tree.getroot(), 'wd', structs, True)

    print(f'// Generated from {file}')
    print(f'package {package}\n')
    print('import "encoding/xml"\n')
    for _, (name, fields) in enumerate(structs.items()):
        print('\n'.join(gen_struct(name, fields)))
        print()

if __name__ == '__main__':
    if len(sys.argv) != 3:
        print("Usage: python generate_go_structs.py <xml_file_glob> <package>")
        sys.exit(1)
    main(sys.argv[1], sys.argv[2])