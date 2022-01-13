# Helm Parser
### Runs a helm parser in the browser via webassembly that runs compiled golang code.
***
<br>I was looking for an online Helm parser and stumbled upon [a repo](https://github.com/radu-matei/helm-template-wasm) via this [great blog post](https://radu-matei.com/blog/helm-wasm/). There was no website but it was possible to clone the repo and from that I have tweaked some stuff and put it online.  

I think Webassembly is an excellent fit for this type of project where you can run the real validation code written in golang directly in your browser.  

You can try it out here: https://helm.blandann.at/
