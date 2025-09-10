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
> with requests coming from `LOCAL_CIDRS`, `LOCAL_IP` will be responsed.
- **response**
    ```json
    { "ip": "your ip" }
    ```