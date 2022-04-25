# xbin-local

script is comming from 'https://github.com/xbin-io/xbin'

the script may changed to :

```
# XBIN_SITE="https://xbin.io"
XBIN_SITE="http://localhost:7890"
function xbin() {
  command="$1"
  args="${@:2}"
  if [ -t 0 ]; then
    curl -X "POST ${XBIN_SITE}/${command}" -H "X-Args: ${args}"
  else
    curl --data-binary @- "${XBIN_SITE}/${command}" -H "X-Args: ${args}"
  fi
}
```

using local server, you could control everything you want.


## build and run

```
go build
./xbin-local
```
