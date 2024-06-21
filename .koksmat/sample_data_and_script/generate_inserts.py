
import yaml

def load_data(yaml_file):
    with open(yaml_file, 'r') as file:
        return yaml.safe_load(file)

def generate_insert_statements(data):
    statements = []

    for country in data['countries']:
        statements.append(f"INSERT INTO country (id, name, code, description) VALUES ({country['id']}, '{country['name']}', '{country['code']}', '{country['description']}');")

    for site in data['sites']:
        statements.append(f"INSERT INTO site (id, name, address, description, country_id) VALUES ({site['id']}, '{site['name']}', '{site['address']}', '{site['description']}', {site['country_id']});")

    for building in data['buildings']:
        statements.append(f"INSERT INTO building (id, name, address, description, site_id) VALUES ({building['id']}, '{building['name']}', '{building['address']}', '{building['description']}', {building['site_id']});")

    for floor in data['floors']:
        statements.append(f"INSERT INTO floor (id, name, building, building_id) VALUES ({floor['id']}, '{floor['name']}', '{floor['building']}', {floor['building_id']});")

    return statements

def main():
    data = load_data('sample_data.yaml')
    statements = generate_insert_statements(data)
    for statement in statements:
        print(statement)

if __name__ == "__main__":
    main()
