# Kiichain with custom module

In this project kiichain has been used for the creation of the blockchain. If you want to know more about kiichain, please consult [kiichain](https://github.com/KiiBlockchain/kii).

The project contains the addition of a custom module that allows CRUD operations on a list of `post`.

You can run the project on your local or using docker (for personal preference for the execution I prefer to use docker).

## RUN

### LOCAL

You can run it on your premises but you must take into account the prerequisites.

- go
- [ignite cli](https://docs.ignite.com/v0.27/welcome/install)

Run to run the project locally:

```bash
ignite chain serve
```

### Docker

I recommend the use of docker since it allows a controlled environment in which you can limit the use of resources for the container, you can have a totally isolated environment if required without having any conflict with any dependencies or installed libraries.

Note that you must have previously installed `docker`.

1. Compile docker image

    ```bash
    # using the common builder
    docker build -t kii . 
    # using builx
    docker buildx build -t kii . 
    ```
2. Run your container with controlled resources: I recommend this because running the blockchain uses all available CPU and RAM resources. By running it in a container we can allocate a certain amount of resources.

    ```bash
    docker run -it \
        --name kii \
        -p 1317:1317 \
        -p 26656:26656 \
        -p 26657:26657 \
        -p 26658:26658 \
        -p 6060:6060 \
        -p 8080:8080 \
        -p 9090:9090 \
        -p 9091:9091 \
        --memory-reservation 11776m \
        --memory 11776m \
        --cpus 4 \
        --workdir /app \
        --rm kii
    ```

3. Execute the project inside the container: when executing the above command by default your terminal is connected to the container terminal so you can execute the following command:

    ```bash
    ignite chain serve
    ```

## Interacting with the project

Regardless of how you run the project you can access the following links:

`swagger`: http://0.0.0.0:1317/  

http://0.0.0.0:26657/

if you want to interact with the blog module you can create a blog using the following command

```bash
kiichaind tx kiichain create-post 'Hello, World!' 'This is a blog post' -
-from exampleWalletKiiBotom
```

if you want to read the saved posts you can do it through the console or with swagger.
