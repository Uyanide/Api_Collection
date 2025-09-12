## API Collection for self use

### General envs

- **PORT**: port that the server listens to (default: 10087)
- **GIN_MODE**: as defined in [Gin Documentation](https://gin-gonic.com/en/docs/deployment/)

### IP query

- **method**: GET
- **route**: /ip
- **envs**:
    - **LOCAL_IP**: public ip of the local system (defaut: 127.0.0.1 <s>which does not make any sense, I know</s>)
    - **LOCAL_CIDRS**: CIDRs of subnets & private networks (defaut: empty). e.g.
        ```conf
        LOCAL_CIDRS="10.0.0.0/8,172.16.0.0/12,192.168.0.0/16,127.0.0.0/8,::1/128,fe80::/10,fc00::/7,fd00::/8"
        ```
> [!NOTE]
>
> For requests coming from `LOCAL_CIDRS`, `LOCAL_IP` will be responded.
- **response**:
    ```json
    { "ip": "your ip" }
    ```

### File service

- **method**: GET
- **route**: as defined in `FILE_MAP`
- **envs**:
    - **FILE_MAP**: entries separated with `,`, each entry consists of three elements sepatated with `:`:
        - url_path: e.g. `/myfile`
        - file_path: e.g. `/app/data/myfile.txt`
        - file_name: e.g. `myfile_downloaded.txt`

        as a whole,
        ```conf
        FILE_MAP=/myfile:/app/data/myfile.txt:myfile_downloaded.txt
        ```
        means `domain.tld/myfile` will download `/app/data/myfile.txt` with name `myfile_downloaded.txt`
- **response**: file as attachment
> [!IMPORTANT]
>
> If the service runs in a docker container, make sure file_path in `FILE_MAP` points to files within the container rather than on the local disk. Consider placing the file in `./data`, which will map to `/app/data` within the container.

### CORS proxy

- **method**: any (not completely tested yet)
- **route**: /proxy
- **params**:
    - **url**: url with http/https scheme
- **envs**:
    - **AUTO_CORRECT_SCHEME**: whether to add `http://` to urls that do not begin with `http://` or `https://`. accepted values: `0` / `1`, default: `0`

### Stats

- **method**: GET
- **route**: /stats
- **responce**:
    ```json
    {
        "total_requests": 0,
        "ip_requests": {
            "total_requests": 0
        },
        "file_downloads": {
            "total_downloads": 0,
            "most_downloaded": "/somefile",
            "files": [
                {
                    "url_path": "/somefile",
                    "downloads_count": 0
                }
            ]
        },
        "proxied_requests": {
            "total_requests": 0,
            "get": 0,
            "post": 0,
            "put": 0,
            "delete": 0
        }
    }
    ```