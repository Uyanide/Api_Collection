## API Collection for self use

### General envs

- **PORT**: port that the server listens to (default: 10087)

### IP query

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
