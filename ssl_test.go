// Copyright (C) 2017. See AUTHORS.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openssl

var (
	certBytes = []byte(`-----BEGIN CERTIFICATE-----
MIIDxDCCAqygAwIBAgIVAMcK/0VWQr2O3MNfJCydqR7oVELcMA0GCSqGSIb3DQEB
BQUAMIGQMUkwRwYDVQQDE0A1NjdjZGRmYzRjOWZiNTYwZTk1M2ZlZjA1N2M0NGFm
MDdiYjc4MDIzODIxYTA5NThiY2RmMGMwNzJhOTdiMThhMQswCQYDVQQGEwJVUzEN
MAsGA1UECBMEVXRhaDEQMA4GA1UEBxMHTWlkdmFsZTEVMBMGA1UEChMMU3BhY2Ug
TW9ua2V5MB4XDTEzMTIxNzE4MzgyMloXDTIzMTIxNTE4MzgyMlowgZAxSTBHBgNV
BAMTQDM4NTg3ODRkMjU1NTdiNTM1MWZmNjRmMmQzMTQ1ZjkwYTJlMTIzMDM4Y2Yz
Mjc1Yzg1OTM1MjcxYWIzMmNiMDkxCzAJBgNVBAYTAlVTMQ0wCwYDVQQIEwRVdGFo
MRAwDgYDVQQHEwdNaWR2YWxlMRUwEwYDVQQKEwxTcGFjZSBNb25rZXkwggEiMA0G
CSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDdf3icNvFsrlrnNLi8SocscqlSbFq+
pEvmhcSoqgDLqebnqu8Ld73HJJ74MGXEgRX8xZT5FinOML31CR6t9E/j3dqV6p+G
fdlFLe3IqtC0/bPVnCDBirBygBI4uCrMq+1VhAxPWclrDo7l9QRYbsExH9lfn+Ry
vxeNMZiOASasvVZNncY8E9usBGRdH17EfDL/TPwXqWOLyxSN5o54GTztjjy9w9CG
QP7jcCueKYyQJQCtEmnwc6P/q6/EPv5R6drBkX6loAPtmCUAkHqxkWOJrRq/v7Pw
zRYhfY+ZpVHGc7WEkDnLzRiUypr1C9oxvLKS10etZEIwEdKyOkSg2fdPAgMBAAGj
EzARMA8GA1UdEwEB/wQFMAMCAQAwDQYJKoZIhvcNAQEFBQADggEBAEcz0RTTJ99l
HTK/zTyfV5VZEhtwqu6bwre/hD7lhI+1ji0DZYGIgCbJLKuZhj+cHn2h5nPhN7zE
M9tc4pn0TgeVS0SVFSe6TGnIFipNogvP17E+vXpDZcW/xn9kPKeVCZc1hlDt1W4Z
5q+ub3aUwuMwYs7bcArtDrumCmciJ3LFyNhebPi4mntb5ooeLFLaujEmVYyrQnpo
tWKC9sMlJmLm4yAso64Sv9KLS2T9ivJBNn0ZtougozBCCTqrqgZVjha+B2yjHe9f
sRkg/uxcJf7wC5Y0BLlp1+aPwdmZD87T3a1uQ1Ij93jmHG+2T9U20MklHAePOl0q
yTqdSPnSH1c=
-----END CERTIFICATE-----
`)
	keyBytes = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA3X94nDbxbK5a5zS4vEqHLHKpUmxavqRL5oXEqKoAy6nm56rv
C3e9xySe+DBlxIEV/MWU+RYpzjC99QkerfRP493aleqfhn3ZRS3tyKrQtP2z1Zwg
wYqwcoASOLgqzKvtVYQMT1nJaw6O5fUEWG7BMR/ZX5/kcr8XjTGYjgEmrL1WTZ3G
PBPbrARkXR9exHwy/0z8F6lji8sUjeaOeBk87Y48vcPQhkD+43ArnimMkCUArRJp
8HOj/6uvxD7+UenawZF+paAD7ZglAJB6sZFjia0av7+z8M0WIX2PmaVRxnO1hJA5
y80YlMqa9QvaMbyyktdHrWRCMBHSsjpEoNn3TwIDAQABAoIBAQCwgp6YzmgCFce3
LBpzYmjqEM3CMzr1ZXRe1gbr6d4Mbu7leyBX4SpJAnP0kIzo1X2yG7ol7XWPLOST
2pqqQWFQ00EX6wsJYEy+hmVRXl5HfU3MUkkAMwd9l3Xt4UWqKPBPD5XHvmN2fvl9
Y4388vXdseXGAGNK1eFs0TMjJuOtDxDyrmJcnxpJ7y/77y/Hb5rUa9DCvj8tkKHg
HmeIwQE0HhIFofj+qCYbqeVyjbPAaYZMrISXb2HmcyULKEOGRbMH24IzInKA0NxV
kdP9qmV8Y2bJ609Fft/y8Vpj31iEdq/OFXyobdVvnXMnaVyAetoaWy7AOTIQ2Cnw
wGbJ/F8BAoGBAN/pCnLQrWREeVMuFjf+MgYgCtRRaQ8EOVvjYcXXi0PhtOMFTAb7
djqhlgmBOFsmeXcb8YRZsF+pNtu1xk5RJOquyKfK8j1rUdAJfoxGHiaUFI2/1i9E
zuXX/Ao0xNRkWMxMKuwYBmmt1fMuVo+1M8UEwFMdHRtgxe+/+eOV1J2PAoGBAP09
7GLOYSYAI1OO3BN/bEVNau6tAxP5YShGmX2Qxy0+ooxHZ1V3D8yo6C0hSg+H+fPT
mjMgGcvaW6K+QyCdHDjgbk2hfdZ+Beq92JApPrH9gMV7MPhwHzgwjzDDio9OFxYY
3vjBQ2yX+9jvz9lkvq2NM3fqFqbsG6Et+5mCc6pBAoGBAI62bxVtEgbladrtdfXs
S6ABzkUzOl362EBL9iZuUnJKqstDtgiBQALwuLuIJA5cwHB9W/t6WuMt7CwveJy0
NW5rRrNDtBAXlgad9o2bp135ZfxO+EoadjCi8B7lMUsaRkq4hWcDjRrQVJxxvXRN
DxkVBSw0Uzf+/0nnN3OqLODbAoGACCY+/isAC1YDzQOS53m5RT2pjEa7C6CB1Ob4
t4a6MiWK25LMq35qXr6swg8JMBjDHWqY0r5ctievvTv8Mwd7SgVG526j+wwRKq2z
U2hQYS/0Peap+8S37Hn7kakpQ1VS/t4MBttJTSxS6XdGLAvG6xTZLCm3UuXUOcqe
ByGgkUECgYEAmop45kRi974g4MPvyLplcE4syb19ifrHj76gPRBi94Cp8jZosY1T
ucCCa4lOGgPtXJ0Qf1c8yq5vh4yqkQjrgUTkr+CFDGR6y4CxmNDQxEMYIajaIiSY
qmgvgyRayemfO2zR0CPgC6wSoGBth+xW6g+WA8y0z76ZSaWpFi8lVM4=
-----END RSA PRIVATE KEY-----
`)
	prime256v1KeyBytes = []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIB/XL0zZSsAu+IQF1AI/nRneabb2S126WFlvvhzmYr1KoAoGCCqGSM49
AwEHoUQDQgAESSFGWwF6W1hoatKGPPorh4+ipyk0FqpiWdiH+4jIiU39qtOeZGSh
1QgSbzfdHxvoYI0FXM+mqE7wec0kIvrrHw==
-----END EC PRIVATE KEY-----
`)
	prime256v1CertBytes = []byte(`-----BEGIN CERTIFICATE-----
MIIChTCCAiqgAwIBAgIJAOQII2LQl4uxMAoGCCqGSM49BAMCMIGcMQswCQYDVQQG
EwJVUzEPMA0GA1UECAwGS2Fuc2FzMRAwDgYDVQQHDAdOb3doZXJlMR8wHQYDVQQK
DBZGYWtlIENlcnRpZmljYXRlcywgSW5jMUkwRwYDVQQDDEBhMWJkZDVmZjg5ZjQy
N2IwZmNiOTdlNDMyZTY5Nzg2NjI2ODJhMWUyNzM4MDhkODE0ZWJiZjY4ODBlYzA3
NDljMB4XDTE3MTIxNTIwNDU1MVoXDTI3MTIxMzIwNDU1MVowgZwxCzAJBgNVBAYT
AlVTMQ8wDQYDVQQIDAZLYW5zYXMxEDAOBgNVBAcMB05vd2hlcmUxHzAdBgNVBAoM
FkZha2UgQ2VydGlmaWNhdGVzLCBJbmMxSTBHBgNVBAMMQGExYmRkNWZmODlmNDI3
YjBmY2I5N2U0MzJlNjk3ODY2MjY4MmExZTI3MzgwOGQ4MTRlYmJmNjg4MGVjMDc0
OWMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARJIUZbAXpbWGhq0oY8+iuHj6Kn
KTQWqmJZ2If7iMiJTf2q055kZKHVCBJvN90fG+hgjQVcz6aoTvB5zSQi+usfo1Mw
UTAdBgNVHQ4EFgQUfRYAFhlGM1wzvusyGrm26Vrbqm4wHwYDVR0jBBgwFoAUfRYA
FhlGM1wzvusyGrm26Vrbqm4wDwYDVR0TAQH/BAUwAwEB/zAKBggqhkjOPQQDAgNJ
ADBGAiEA6PWNjm4B6zs3Wcha9qyDdfo1ILhHfk9rZEAGrnfyc2UCIQD1IDVJUkI4
J/QVoOtP5DOdRPs/3XFy0Bk0qH+Uj5D7LQ==
-----END CERTIFICATE-----
`)
	ed25519CertBytes = []byte(`-----BEGIN CERTIFICATE-----
MIIBIzCB1gIUd0UUPX+qHrSKSVN9V/A3F1Eeti4wBQYDK2VwMDYxCzAJBgNVBAYT
AnVzMQ0wCwYDVQQKDARDU0NPMRgwFgYDVQQDDA9lZDI1NTE5X3Jvb3RfY2EwHhcN
MTgwODE3MDMzNzQ4WhcNMjgwODE0MDMzNzQ4WjAzMQswCQYDVQQGEwJ1czENMAsG
A1UECgwEQ1NDTzEVMBMGA1UEAwwMZWQyNTUxOV9sZWFmMCowBQYDK2VwAyEAKZZJ
zzlBcpjdbvzV0BRoaSiJKxbU6GnFeAELA0cHWR0wBQYDK2VwA0EAbfUJ7L7v3GDq
Gv7R90wQ/OKAc+o0q9eOrD6KRYDBhvlnMKqTMRVucnHXfrd5Rhmf4yHTvFTOhwmO
t/hpmISAAA==
-----END CERTIFICATE-----
`)
	ed25519KeyBytes = []byte(`-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIL3QVwyuusKuLgZwZn356UHk9u1REGHbNTLtFMPKNQSb
-----END PRIVATE KEY-----
`)
)
