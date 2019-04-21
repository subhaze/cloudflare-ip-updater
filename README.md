# cloudflare-ip-updater
Used as a Dynamic DNS updater for Cloudflare on Synology NAS

## Environment variables needed

- `CF_API_KEY`
- `CF_API_EMAIL`

## Usage

You can start it with:<br>`docker run --rm -e CF_API_KEY=YOUR_API_KEY -e CF_API_EMAIL=YOUR_CF_EMAIL -p 8080:8080 subhaze/cloudflare-ip-updater`

Navigate to `http://localhost:8080/cf-ip-update?site=SITE&domain=DOMAIN&ip=IP_YOU_WANT_TO_UPDATE`

`site`   name of the site in Cloudflare, such as `example.com`<br>
`domain` the domain you wish to update in the DNS, this could be `example.com`, `test.example.com`, and so on.
