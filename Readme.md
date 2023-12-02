# JSON Schema to JavaScript Class Converter

This tool is a simple and effective utility written in Go, designed to convert JSON Schemas into JavaScript classes. It supports fetching JSON Schemas from either local files or remote URLs and outputs JavaScript classes with ES6 syntax.

## Features

- Fetch JSON Schema from a local file or a remote URL.
- Generate JavaScript classes with getters and setters.
- Output the generated JavaScript classes to specified directories.
- Easy to integrate into existing workflows.

## Getting Started

### Prerequisites

Ensure you have [Go](https://golang.org/dl/) installed on your machine. This tool is compatible with Go version 1.15 or higher.

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/your-username/json-schema-to-js.git
cd json-schema-to-js
```

### Usage`

Run the tool using the Go command:

```bash
go run main.go -file <path-to-json-schema> -o <output-directory>
```

or

```bash
go run main.go -url <url-of-json-schema> -o <output-directory>
```

#### Flags

- `-file`: Path to the local JSON Schema file.
- `-url`: URL of the JSON Schema.
- `-o`: (Optional) Path to the output directory for the generated JavaScript classes. If omitted, the output will be printed to the standard output.

### Example

#### person.json
```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/person.schema.json",
  "$ref": "#/$defs/Person",
  "$defs": {
    "Person": {
      "type": "object",
      "properties": {
        "firstName": {
          "type": "string"
        },
        "lastName": {
          "type": "string"
        },
        "age": {
          "type": "integer",
          "minimum": 0
        },
        "email": {
          "type": "string",
          "format": "email"
        },
        "address": {
          "$ref": "#/$defs/Address"
        }
      },
      "required": ["firstName", "lastName", "age", "address"]
    },
    "Address": {
      "type": "object",
      "properties": {
        "streetAddress": {
          "type": "string"
        },
        "city": {
          "type": "string"
        },
        "state": {
          "type": "string"
        },
        "postalCode": {
          "type": "string"
        }
      },
      "required": ["streetAddress", "city", "state", "postalCode"]
    }
  }
}
```

```bash
go run main.go -file ./person.json -o ./output
```

this will generate

#### Person.js
```javascript
export class Person {
  constructor() {
    this._FirstName = "";
    this._LastName = "";
    this._Age = "";
    this._Address = "";
  }

  static fromJson(json) {
    let person = new Person();
    person.FirstName = json.firstName;
    person.LastName = json.lastName;
    person.Age = json.age;
    person.Address = json.address;
    return person;
  }

  get FirstName() {
    return this._FirstName;
  }

  set FirstName(value) {
    this._FirstName = value;
  }

  get LastName() {
    return this._LastName;
  }

  set LastName(value) {
    this._LastName = value;
  }

  get Age() {
    return this._Age;
  }

  set Age(value) {
    this._Age = value;
  }

  get Address() {
    return this._Address;
  }

  set Address(value) {
    this._Address = value;
  }
}
```

and

#### Address.js
```javascript
import {Address} from "./Address"
export class Person {
  constructor() {
    this._FirstName = "";
    this._LastName = "";
    this._Age = "";
    this._Address = "";
  }

  static fromJson(json) {
    let person = new Person();
    person.Address = Address.fromJson(json.address);
    person.FirstName = json.firstName;
    person.LastName = json.lastName;
    person.Age = json.age;
    person.Email = json.email;
    return person;
  }

  get FirstName() {
    return this._FirstName;
  }

  set FirstName(value) {
    this._FirstName = value;
  }

  get LastName() {
    return this._LastName;
  }

  set LastName(value) {
    this._LastName = value;
  }

  get Age() {
    return this._Age;
  }

  set Age(value) {
    this._Age = value;
  }

  get Address() {
    return this._Address;
  }

  set Address(value) {
    this._Address = value;
  }
}
```

This command will read the JSON Schema from `./schemas/customer.json` and generate JavaScript classes in the `./output` directory.

## Contributing

Contributions to enhance this tool are welcome. Please fork the repository and submit a pull request with your improvements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the Go community for providing the essential libraries used in this project.
- This tool was created to simplify the workflow of converting JSON Schema to JavaScript, and we hope it helps others in their development processes.
