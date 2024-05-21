# Email username generator

Generates random usernames for use in email addresses.

## Config

You can use following environment variables to configure this service:

- `EPG_PREFIX`: Prefix of the username. Defaults to `ext`.
- `EPG_SEPARATOR`: Separator between prefix, external party name and suffix. Defaults to `.`
- `EPG_SUFFIX_RANDOM_SET`: Set of characters to use for suffix. Defaults to lowercase alphanumeric.

Service is available on port `8080`.
