# httpcache
> 💾 Cache HTTP request easily from public API - https://cache.smiti.fr
# Goal
HttpCache is a http wrapper that allows you to cache
http request easily without configuration.
It can be very useful for cache public API request
on a static website for example, not needed anymore
to develop and deploy a http proxy to do this job.
# Usage
you just need to add `https://cache.smiti.fr/{exp_in_minute}/` before the url of the service you want cache, that's it !
You can see an example below :
```shell
https://cache.smiti.fr/1280/https://api.github.com/users/ermos/repos?sort=created&per_page=4&page=1&desc
```
# Use cases

## Cache request across users

### Without HTTPCache.me
<p align="center">
  <img src="docs/without_httpcache.jpg">
</p>
When you use a public API directly in your website,
the request is made from the user's client, if the public API send cache header,
the user's client will save it. Now, if a second user come on your website,
he will send a new request and store it into his client.
Each request need working server side and can cost time.

### With HTTPCache.me
<p align="center">
  <img src="docs/with_httpcache.jpg">
</p>
When a user come to your website, the request is made from the user's client, httpcache will save the result in memory and when
the second user will come on your website, the result will directly returned without communicate with the public API.
