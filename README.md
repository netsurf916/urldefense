# urldefense
A tool to create a local webserver which can take a "urldefense.com" URI and send a redirect to the destination in order to bypass sandboxing (which is generally slow, buggy, or both).

In order to make use of this tool, you must execute "setup.sh" to create local certificates -- which will need to be trusted on the first connection.  You must also redirect urldefense.com to localhost via the hosts file or other means.
