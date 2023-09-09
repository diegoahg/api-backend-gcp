# api-backend-gcp
Api Backend subida al artifactory de GCP

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)
## Installation
Pre-require:
 - go 1.21.0 or higher

```bash
# Ejemplo de comandos de instalación
git clone https://github.com/diegoahg/api-backend-gcp.git
cd api-backend-gcp
go build -o api-backend
```
# Usage
if you have linux or MacOs
```bash
./api-backend
```
for windows
```bash
./api-backend.exe
```

# Testing
go to cmd folder 
```bash
go cd cmd
```
```bash
go test ./...
```
Something this appears:
```bash
?       api-backend-gcp/api     [no test files]
?       api-backend-gcp/cmd     [no test files]
ok      api-backend-gcp/pkg/palabraazar 0.195s  coverage: 100.0% of statements
```

## Project Structure
api-backend-gcp/
    ├── cmd/
    │   └── main.go
    ├── pkg/
    │   ├── palabraazar/
    │   │   ├── palabraazar.go
    │   │   └── palabraazar_test.go
    ├── api/
    │   ├── api.go
    │   └── api_test.go
    ├── go.mod
    ├── go.sum

## Contributing
Thank you for considering contributing to our project! If you'd like to contribute, fork and submit a pull request!

## License
MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
