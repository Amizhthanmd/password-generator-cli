# Password Generator CLI

A simple command-line tool to generate secure passwords with customizable options. Built with Go.

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/Amizhthanmd/password-generator-cli.git

cd password-generator-cli

go build -o password-generator
```

### Usage

Generate a password with various customizable options:

### Basic Command

```bash
./password-generator generate
```

# Example

Generate a password with a specific length and no uppercase letters:

```bash
./password-generator generate -t 16 -u=false
```

Generate a password without lowercase letters:

```bash
./password-generator generate -l=false
```

Generate a password without symbols:

```bash
./password-generator generate -s=false
```
