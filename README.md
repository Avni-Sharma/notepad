# notepad

## Build binary

```sh
go build
```

## Build & Push image

```sh
docker login quay.io -u <username>
docker build -t quay.io/<username>/appbinding-demo-notepad .
docker push quay.io/<username>/appbinding-demo-notepad
```

## Deploy

```sh
sed -e 's,REPLACE_IMAGE,quay.io/<username>/appbinding-demo-notepad,g' deploy.yaml | oc apply -f -
oc apply -f service.yaml
oc apply -f route.yaml
```

## Use notepad application

### Discover notepad endpoint

```sh
export NOTEPAD_ENDPOINT="$(oc get route notepad-route -o jsonpath='{.spec.host}')/note"
```

### Get content

```sh
curl -XGET "$NOTEPAD_ENDPOINT"
```

### Append content

```sh
curl -XPOST -H "ContentType=text/plain" -d "text to append" "$NOTEPAD_ENDPOINT"
```
