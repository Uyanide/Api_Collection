## A tiny API collection for self use and practice

### Usage

Self-hosted deployment:

1. Clone this repository;
2. Create and edit `docker/.env`. All possible envs are listed below;
3. Execute `docker/run.sh`.
> [!IMPORTANT]
>
> `PORT` env is required when deploying with Docker. Either set it in your shell or add it to `docker/.env`.

Develop or deploy without Docker:

1. Clone this repository;
2. Create and edit `.env`. All possible envs are listed below;
3. `go run .` or other common methods to build and run a Go program.

>Example `docker/.env`:
>```shell
>PORT="10087"
>GIN_MODE="release"
>DB_PATH="/app/data/db"
>LOCAL_IP="19.19.8.10"
>LOCAL_CIDRS="10.0.0.0/8,172.16.0.0/12,192.168.0.0/16,127.0.0.0/8,::1/128,fe80::/10,fc00::/7,fd00::/8"
>FILE_MAP="/hello:/app/data/hello.txt:hello.txt"
>AUTO_CORRECT_SCHEME="0"
>```

### General envs

- **PORT**: port that the server listens to (default: 10087)
- **GIN_MODE**: as defined in [Gin Documentation](https://gin-gonic.com/en/docs/deployment/)
- **DB_PATH**: path to database directory (default `data/db`)

### IP query

- **method**: GET
- **route**: /ip
- **envs**:
    - **LOCAL_IP**: public ip of the local system (defaut: 127.0.0.1 <s>which does not make any sense, I know</s>)
    - **LOCAL_CIDRS**: CIDRs of subnets & private networks (defaut: empty). e.g.
        ```shell
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
        ```shell
        FILE_MAP="/myfile:/app/data/myfile.txt:myfile_downloaded.txt"
        ```
        means `/app/data/myfile.txt` can be downloaded via `http://domain.tld/myfile` with name `myfile_downloaded.txt`
- **response**: file as attachment
> [!IMPORTANT]
>
> If the service runs in a Docker container, make sure file_path in `FILE_MAP` points to files within the container rather than on the local disk. Consider placing the file in `./data`, which will map to `/app/data` within the container.

### File service (in directories)

- **method**: GET
- **route**: `${url_path}/*filepath`, where `url_path` is defined in `DIR_MAP`
- **envs**:
    - **DIR_MAP**: entries separated with `,`, each entry consists of two elements sepatated with `:`:
        - url_path: e.g. `/wallpapers`
        - dir_path: e.g. `/app/data/backgrounds`

        for example, if `DIR_MAP` is set to
        ```shell
        DIR_MAP="/wallpapers:/app/data/wallpapers"
        ```
        and there is a file `/app/data/backgrounds/nature/mountain.jpg`, then it can be downloaded via `http://domain.tld/wallpapers/nature/mountain.jpg`, and the file name will be `mountain.jpg`.
- **response**:
    - if the path points to a file, the file will be sent as attachment;
    - if the path points to a directory, an HTML page listing all files and sub-directories will be sent.

> [!IMPORTANT]
>
> If the service runs in a Docker container, make sure path in `DIR_MAP` points to dir within the container rather than on the local disk. Consider placing the files in `./data`, which will map to `/app/data` within the container.

> [!CAUTION]
>
> No trailing slash is allowed in `url_path` and `dir_path` in `DIR_MAP`.



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
            "successful": 0,
            "get": 0,
    }
    ```
