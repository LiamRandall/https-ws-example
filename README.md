# https-ws-example

## Overview
This toy example demonstrates the typical security headers you might see running on a website in production. These included example spins up a simple webserver on 8080 and websocket server on 8081.  The webserver on 8080 makes writes HTML to instruct your browser to make calls to the websocket server on 8081.

## Risks
When front-end applications use WebSockets and HTTP connections on different domains, several security concerns and errors can arise:

#### Cross-Origin Resource Sharing (CORS) Issues: 
Browsers implement the same-origin policy, which restricts how a document or script loaded from one origin can interact with resources from another origin. If your WebSocket and HTTP connections are on different domains, you might encounter CORS errors unless proper headers are set to allow such requests.

#### Mixed Content Warnings: 
If your main application is served over HTTPS but the WebSocket connection uses an unsecured (ws://) connection, browsers will generate mixed content warnings. This can lead to blocking of unsecured content on secure pages.

#### Cookie Handling Issues: 
Cookies are often used for maintaining session state. If your HTTP and WebSocket connections are on different domains, managing authentication and session cookies can become complicated, leading to security risks if not handled properly.

#### Cross-Site WebSocket Hijacking (CSWSH): 
Similar to Cross-Site Request Forgery (CSRF), if the WebSocket server does not properly validate the origin of the connection, an attacker could establish a WebSocket connection from a malicious site, potentially leading to unauthorized actions.

#### Security Policy Configuration Complexity: 
Having different domains for HTTP and WebSocket connections increases the complexity of configuring security policies like Content Security Policy (CSP), which can lead to misconfigurations and vulnerabilities.

#### SSL/TLS Certificate Management: 
If you're using HTTPS for your main application and WSS (WebSocket Secure) for WebSockets, you'll need to manage SSL/TLS certificates for both domains, which increases complexity and the risk of misconfiguration.

#### Domain-Based Security Controls: 
Security mechanisms that rely on domain reputation (like some firewalls or intrusion detection systems) might treat traffic differently based on the domain, potentially leading to blocked or flagged connections.

#### Subdomain Takeover Risks: 
If you're using a subdomain for one of the services (e.g., WebSocket), there's a risk of subdomain takeover if the domain expires or is not properly managed.


### Explanation of Security Headers Demonstrated
#### Content-Security-Policy (CSP): 
Restricts resources the client is allowed to load for a given page. This helps mitigate XSS attacks.

#### X-Content-Type-Options: 
Prevents the browser from interpreting files as a different MIME type than what is specified in the Content-Type HTTP header.

#### X-Frame-Options: 
Protects against clickjacking attacks by preventing your webpage from being embedded in an iframe.

#### X-XSS-Protection: 
Enables the browser's built-in XSS filter.

#### Strict-Transport-Security (HSTS): 
Enforces secure (HTTP over SSL/TLS) connections to the server. This header is commented out because it only works over HTTPS.

#### Referrer-Policy: 
Controls how much referrer information should be included with requests.

#### Feature-Policy: 
Allows you to selectively enable, disable, or modify the behavior of certain APIs and web features in the browser.

