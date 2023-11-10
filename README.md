![Static Badge](https://img.shields.io/badge/-Language-blue?logo=Go&label=1.20)
![GitHub](https://img.shields.io/github/license/embe-codez/domain-detective?logo=GitHub&color=blue)

# Domain Detective

Domain Detective is a command-line tool developed in Go (Golang) that allows users to perform various domain-related operations. It provides insights into a domain's DNS records, WHOIS data, and SSL certificate details.

## Screenshot
![image](https://github.com/Embe-Codez/DomainDetective/assets/74840238/5b26b7db-95a7-40d0-b55c-a35d7edfd73e)

## Features

* DNS Lookup: Retrieve DNS records (A, CNAME, TXT) for a domain.
* WHOIS Lookup: Get WHOIS information for a domain.
* Certificate Information: Fetch SSL certificate details for a domain.

## Usage

1. Clone the repository: ```git clone https://github.com/Embe-Codez/domain-detective.git```

2. Build the executable: ```go build```

3. Run the tool: ```./DomainDetective```

4. Follow the prompts to enter the domain name and record type.

## Dependencies

The project has the following external dependency:

* This tool requires you to install Go locally on your machine, which can be done by following the instructions provided at ```https://go.dev/doc/install```.

* Go package for performing WHOIS lookups - ```github.com/undiabler/golang-whois``` 

## Contributing

Contributions to the Domain Detective project are welcome! If you find any bugs or have suggestions for improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for more information.
Author

The Domain Detective project is developed and maintained by [Embe-Codez].
