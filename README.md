# Minaris
My wife can do any fukin' thing.

## Shooter mode

shoot:
```
minaris shooter -f [workflow file]
```

Workflow file example:
```yaml
name: example workflow
workflows:
  - name: Load facts
    # run the "fact" module 
    fact:
      raw:
        - KEY=VALUE

  - name: Login
    # run the "shoot" module
    shoot:
      file: ./vanieva/login.k4it0z11
```

Raw flow file:
```
label_1
====minaris====
[[string selected]]--->"user":"0938194095"
"captchaValue":"zvdg8"--->"captchaValue":"3gke3"
"captchaToken":"f86b9833-26fa-1c26-6c57-a78486c47d24"--->"captchaToken":"ad7cc79d-18d1-7aa8-1135-49e248d01751"
====minaris====
LOGIN_TOKEN="token":"([a-z0-9\-]{36})"
USER_PHONE="user":"(\d+)"
====minaris====
[[raw request]]
====k4it0z11====
label_2
====minaris====
[[string_befor]]--->"user":"{{USER_PHONE}}"
"token":"2dcbc544-652d-4c41-8ca6-c3b25341f21c"--->"token":"{{LOGIN_TOKEN}}"
====minaris====
VARNAME=[[regex]]
====minaris====
[[raw request]]
```