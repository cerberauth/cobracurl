package cobracurl

import (
	"github.com/spf13/pflag"
)

func RegisterConnectTimeoutFlag(flags *pflag.FlagSet) {
	flags.Float64("connect-timeout", 0, "Maximum time allowed for connection")
}

func RegisterMaxTimeFlag(flags *pflag.FlagSet) {
	flags.Float64P("max-time", "m", 0, "Maximum time allowed for the transfer")
}

func RegisterProxyFlag(flags *pflag.FlagSet) {
	flags.StringP("proxy", "x", "", "Use this proxy")
}

func RegisterNoproxyFlag(flags *pflag.FlagSet) {
	flags.String("noproxy", "", "List of hosts which do not use proxy")
}

func RegisterProxyAnyauthFlag(flags *pflag.FlagSet) {
	flags.Bool("proxy-anyauth", false, "Pick any proxy authentication method")
}

func RegisterProxyBasicFlag(flags *pflag.FlagSet) {
	flags.Bool("proxy-basic", false, "Use Basic authentication on the proxy")
}

func RegisterProxyCacertFlag(flags *pflag.FlagSet) {
	flags.String("proxy-cacert", "", "CA certificate to verify peer against for proxy")
}

func RegisterProxyCapathFlag(flags *pflag.FlagSet) {
	flags.String("proxy-capath", "", "CA directory to verify peer against for proxy")
}

func RegisterProxyCertFlag(flags *pflag.FlagSet) {
	flags.String("proxy-cert", "", "Set client certificate for proxy")
}

func RegisterProxyCertTypeFlag(flags *pflag.FlagSet) {
	flags.String("proxy-cert-type", "", "Client certificate type for proxy")
}

func RegisterProxyCiphersFlag(flags *pflag.FlagSet) {
	flags.String("proxy-ciphers", "", "SSL ciphers to use for proxy")
}

func RegisterProxyCrlfileFlag(flags *pflag.FlagSet) {
	flags.String("proxy-crlfile", "", "Set a CRL file for proxy")
}

func RegisterProxyDigestFlag(flags *pflag.FlagSet) {
	flags.Bool("proxy-digest", false, "Use Digest authentication on the proxy")
}

func RegisterProxyHeaderFlag(flags *pflag.FlagSet) {
	flags.String("proxy-header", "", "Pass custom header to proxy")
}

func RegisterProxyHttp2Flag(flags *pflag.FlagSet) {
	flags.Bool("proxy-http2", false, "Use HTTP/2 for proxy connections")
}

func RegisterProxyInsecureFlag(flags *pflag.FlagSet) {
	flags.Bool("proxy-insecure", false, "Do https proxy connections without verifying the proxy")
}

func RegisterProxyKeyFlag(flags *pflag.FlagSet) {
	flags.String("proxy-key", "", "Private key for proxy")
}

func RegisterProxyKeyTypeFlag(flags *pflag.FlagSet) {
	flags.String("proxy-key-type", "", "Private key file type for proxy")
}

func RegisterProxyNegotiateFlag(flags *pflag.FlagSet) {
	flags.Bool("proxy-negotiate", false, "Use HTTP Negotiate (SPNEGO) authentication on the proxy")
}

func RegisterProxyNtlmFlag(flags *pflag.FlagSet) {
	flags.Bool("proxy-ntlm", false, "Use NTLM authentication on the proxy")
}

func RegisterProxyPassFlag(flags *pflag.FlagSet) {
	flags.String("proxy-pass", "", "Pass phrase for the private key for proxy")
}

func RegisterProxyServiceNameFlag(flags *pflag.FlagSet) {
	flags.String("proxy-service-name", "", "SPNEGO proxy service name")
}

func RegisterProxySslAllowBeastFlag(flags *pflag.FlagSet) {
	flags.Bool("proxy-ssl-allow-beast", false, "Allow security flaw for interoperability with older servers")
}

func RegisterProxyTls13CiphersFlag(flags *pflag.FlagSet) {
	flags.String("proxy-tls13-ciphers", "", "TLS 1.3 proxy ciphersuites")
}

func RegisterProxyTlsauthtypeFlag(flags *pflag.FlagSet) {
	flags.String("proxy-tlsauthtype", "", "TLS authentication type for proxy")
}

func RegisterProxyTlspasswordFlag(flags *pflag.FlagSet) {
	flags.String("proxy-tlspassword", "", "TLS password for proxy")
}

func RegisterProxyTlsuserFlag(flags *pflag.FlagSet) {
	flags.String("proxy-tlsuser", "", "TLS username for proxy")
}

func RegisterProxyTlsv1Flag(flags *pflag.FlagSet) {
	flags.Bool("proxy-tlsv1", false, "Use TLSv1 for proxy")
}

func RegisterProxyUserFlag(flags *pflag.FlagSet) {
	flags.StringP("proxy-user", "U", "", "Proxy user and password")
}

func RegisterProxy10Flag(flags *pflag.FlagSet) {
	flags.String("proxy1.0", "", "Use HTTP/1.0 proxy on given port")
}

func RegisterProxytunnelFlag(flags *pflag.FlagSet) {
	flags.BoolP("proxytunnel", "p", false, "Tunnel through proxy")
}

func RegisterInsecureFlag(flags *pflag.FlagSet) {
	flags.BoolP("insecure", "k", false, "Allow insecure server connections when using SSL")
}

func RegisterCacertFlag(flags *pflag.FlagSet) {
	flags.String("cacert", "", "CA certificate to verify peer against")
}

func RegisterCapathFlag(flags *pflag.FlagSet) {
	flags.String("capath", "", "CA directory to verify peer against")
}

func RegisterCaNativeFlag(flags *pflag.FlagSet) {
	flags.Bool("ca-native", false, "Use native CA store")
}

func RegisterCertFlag(flags *pflag.FlagSet) {
	flags.StringP("cert", "E", "", "Client certificate file")
}

func RegisterCertStatusFlag(flags *pflag.FlagSet) {
	flags.Bool("cert-status", false, "Verify the status of the server certificate")
}

func RegisterCertTypeFlag(flags *pflag.FlagSet) {
	flags.String("cert-type", "", "Certificate file type (DER/PEM/ENG/P12)")
}

func RegisterCiphersFlag(flags *pflag.FlagSet) {
	flags.String("ciphers", "", "SSL ciphers to use")
}

func RegisterCurvesFlag(flags *pflag.FlagSet) {
	flags.String("curves", "", "SSL/TLS key exchange algorithm(s) to request")
}

func RegisterFalseStartFlag(flags *pflag.FlagSet) {
	flags.Bool("false-start", false, "Enable TLS False Start")
}

func RegisterKeyFlag(flags *pflag.FlagSet) {
	flags.String("key", "", "Private key file name")
}

func RegisterKeyTypeFlag(flags *pflag.FlagSet) {
	flags.String("key-type", "", "Private key file type (DER/PEM/ENG)")
}

func RegisterNoAlpnFlag(flags *pflag.FlagSet) {
	flags.Bool("no-alpn", false, "Disable the ALPN TLS extension")
}

func RegisterNoSessionidFlag(flags *pflag.FlagSet) {
	flags.Bool("no-sessionid", false, "Disable SSL session-ID reusing")
}

func RegisterPassFlag(flags *pflag.FlagSet) {
	flags.String("pass", "", "Pass phrase for the private key")
}

func RegisterPinnedpubkeyFlag(flags *pflag.FlagSet) {
	flags.String("pinnedpubkey", "", "FILE/HASHES public key to verify peer against")
}

func RegisterSslAllowBeastFlag(flags *pflag.FlagSet) {
	flags.Bool("ssl-allow-beast", false, "Allow security flaw for interoperability with older servers")
}

func RegisterTlsMaxFlag(flags *pflag.FlagSet) {
	flags.String("tls-max", "", "Set maximum allowed TLS version")
}

func RegisterTls13CiphersFlag(flags *pflag.FlagSet) {
	flags.String("tls13-ciphers", "", "TLS 1.3 ciphersuites")
}

func RegisterTlsauthtypeFlag(flags *pflag.FlagSet) {
	flags.String("tlsauthtype", "", "TLS authentication type")
}

func RegisterTlspasswordFlag(flags *pflag.FlagSet) {
	flags.String("tlspassword", "", "TLS password")
}

func RegisterTlsuserFlag(flags *pflag.FlagSet) {
	flags.String("tlsuser", "", "TLS user name")
}

func RegisterTlsv1Flag(flags *pflag.FlagSet) {
	flags.BoolP("tlsv1", "1", false, "Use TLSv1.0 or greater")
}

func RegisterTlsv10Flag(flags *pflag.FlagSet) {
	flags.Bool("tlsv1.0", false, "Use TLSv1.0 or greater")
}

func RegisterTlsv11Flag(flags *pflag.FlagSet) {
	flags.Bool("tlsv1.1", false, "Use TLSv1.1 or greater")
}

func RegisterTlsv12Flag(flags *pflag.FlagSet) {
	flags.Bool("tlsv1.2", false, "Use TLSv1.2 or greater")
}

func RegisterTlsv13Flag(flags *pflag.FlagSet) {
	flags.Bool("tlsv1.3", false, "Use TLSv1.3 or greater")
}

func RegisterLocationFlag(flags *pflag.FlagSet) {
	flags.BoolP("location", "L", false, "Follow redirects")
}

func RegisterLocationTrustedFlag(flags *pflag.FlagSet) {
	flags.Bool("location-trusted", false, "Follow redirects, trust all hosts")
}

func RegisterMaxRedirsFlag(flags *pflag.FlagSet) {
	flags.Int("max-redirs", 0, "Maximum number of redirects allowed")
}

func RegisterPost301Flag(flags *pflag.FlagSet) {
	flags.Bool("post301", false, "Do not switch to GET after following a 301")
}

func RegisterPost302Flag(flags *pflag.FlagSet) {
	flags.Bool("post302", false, "Do not switch to GET after following a 302")
}

func RegisterPost303Flag(flags *pflag.FlagSet) {
	flags.Bool("post303", false, "Do not switch to GET after following a 303")
}

func RegisterConnectToFlag(flags *pflag.FlagSet) {
	flags.String("connect-to", "", "Connect to host")
}

func RegisterDohInsecureFlag(flags *pflag.FlagSet) {
	flags.Bool("doh-insecure", false, "Allow insecure DoH server connections")
}

func RegisterDohUrlFlag(flags *pflag.FlagSet) {
	flags.String("doh-url", "", "Resolve host names over DOH")
}

func RegisterHappyEyeballsTimeoutMsFlag(flags *pflag.FlagSet) {
	flags.Int("happy-eyeballs-timeout-ms", 0, "How long to wait for IPv6 before using IPv4")
}

func RegisterInterfaceFlag(flags *pflag.FlagSet) {
	flags.String("interface", "", "Use network INTERFACE (or address)")
}

func RegisterIpv4Flag(flags *pflag.FlagSet) {
	flags.BoolP("ipv4", "4", false, "Resolve to IPv4 addresses")
}

func RegisterIpv6Flag(flags *pflag.FlagSet) {
	flags.BoolP("ipv6", "6", false, "Resolve to IPv6 addresses")
}

func RegisterLocalPortFlag(flags *pflag.FlagSet) {
	flags.String("local-port", "", "Force use of Range of local port numbers")
}

func RegisterResolveFlag(flags *pflag.FlagSet) {
	flags.String("resolve", "", "Resolve the host+port to this address")
}

func RegisterExpect100TimeoutFlag(flags *pflag.FlagSet) {
	flags.Int("expect100-timeout", 0, "How long to wait for 100-continue")
}

func RegisterKeepaliveTimeFlag(flags *pflag.FlagSet) {
	flags.Int("keepalive-time", 0, "Wait seconds before sending keepalive probes")
}

func RegisterNoKeepaliveFlag(flags *pflag.FlagSet) {
	flags.Bool("no-keepalive", false, "Disable TCP keepalive on the connection")
}

func RegisterTcpFastopenFlag(flags *pflag.FlagSet) {
	flags.Bool("tcp-fastopen", false, "Use TCP Fast Open")
}

func RegisterTcpNodelayFlag(flags *pflag.FlagSet) {
	flags.Bool("tcp-nodelay", false, "Use the TCP_NODELAY option")
}

func RegisterAnyauthFlag(flags *pflag.FlagSet) {
	flags.Bool("anyauth", false, "Pick any auth method")
}

func RegisterAwsSigv4Flag(flags *pflag.FlagSet) {
	flags.String("aws-sigv4", "", "Use AWS V4 signature authentication")
}

func RegisterBasicFlag(flags *pflag.FlagSet) {
	flags.Bool("basic", false, "Use HTTP Basic Auth (use with -u)")
}

func RegisterDelegationFlag(flags *pflag.FlagSet) {
	flags.String("delegation", "", "GSS-API delegation permission")
}

func RegisterDigestFlag(flags *pflag.FlagSet) {
	flags.Bool("digest", false, "Use HTTP Digest Auth (use with -u)")
}

func RegisterKrbFlag(flags *pflag.FlagSet) {
	flags.String("krb", "", "Enable Kerberos with security <level>")
}

func RegisterNegotiateFlag(flags *pflag.FlagSet) {
	flags.Bool("negotiate", false, "Use SPNEGO (Negotiate) authentication")
}

func RegisterNetrcFlag(flags *pflag.FlagSet) {
	flags.Bool("netrc", false, "Must read .netrc for user name and password")
}

func RegisterNetrcFileFlag(flags *pflag.FlagSet) {
	flags.String("netrc-file", "", "Specify FILE for netrc")
}

func RegisterNetrcOptionalFlag(flags *pflag.FlagSet) {
	flags.Bool("netrc-optional", false, "Use either .netrc or URL")
}

func RegisterNtlmFlag(flags *pflag.FlagSet) {
	flags.Bool("ntlm", false, "Use NTLM authentication")
}

func RegisterNtlmWbFlag(flags *pflag.FlagSet) {
	flags.Bool("ntlm-wb", false, "Use NTLM authentication with winbind")
}

func RegisterOauth2BearerFlag(flags *pflag.FlagSet) {
	flags.String("oauth2-bearer", "", "OAuth 2 Bearer Token")
}

func RegisterSaslIrFlag(flags *pflag.FlagSet) {
	flags.Bool("sasl-ir", false, "Enable initial response in SASL authentication")
}

func RegisterServiceNameFlag(flags *pflag.FlagSet) {
	flags.String("service-name", "", "SPNEGO service name")
}

func RegisterUserFlag(flags *pflag.FlagSet) {
	flags.StringP("user", "u", "", "Server user and password")
}

func RegisterGetFlag(flags *pflag.FlagSet) {
	flags.BoolP("get", "G", false, "Put the post data in the URL and use GET")
}

func RegisterGloboffFlag(flags *pflag.FlagSet) {
	flags.BoolP("globoff", "g", false, "Disable URL sequences and ranges")
}

func RegisterHeadFlag(flags *pflag.FlagSet) {
	flags.BoolP("head", "I", false, "Show document info only")
}

func RegisterHttp09Flag(flags *pflag.FlagSet) {
	flags.Bool("http0.9", false, "Allow HTTP 0.9 responses")
}

func RegisterHttp10Flag(flags *pflag.FlagSet) {
	flags.BoolP("http1.0", "0", false, "Use HTTP 1.0")
}

func RegisterHttp11Flag(flags *pflag.FlagSet) {
	flags.Bool("http1.1", false, "Use HTTP 1.1")
}

func RegisterHttp2Flag(flags *pflag.FlagSet) {
	flags.Bool("http2", false, "Use HTTP 2")
}

func RegisterHttp2PriorKnowledgeFlag(flags *pflag.FlagSet) {
	flags.Bool("http2-prior-knowledge", false, "Use HTTP 2 without HTTP/1.1 Upgrade")
}

func RegisterHttp3Flag(flags *pflag.FlagSet) {
	flags.Bool("http3", false, "Use HTTP v3")
}

func RegisterHttp3OnlyFlag(flags *pflag.FlagSet) {
	flags.Bool("http3-only", false, "Use HTTP v3 only")
}

func RegisterPathAsIsFlag(flags *pflag.FlagSet) {
	flags.Bool("path-as-is", false, "Do not squash .. sequences in URL path")
}

func RegisterRawFlag(flags *pflag.FlagSet) {
	flags.Bool("raw", false, "Do HTTP 'raw'; no transfer decoding")
}

func RegisterRequestFlag(flags *pflag.FlagSet) {
	flags.StringP("request", "X", "", "Specify request method to use")
}

func RegisterRequestTargetFlag(flags *pflag.FlagSet) {
	flags.String("request-target", "", "Specify the target for this request")
}

func RegisterUrlFlag(flags *pflag.FlagSet) {
	flags.String("url", "", "URL to work with")
}

func RegisterCompressedFlag(flags *pflag.FlagSet) {
	flags.Bool("compressed", false, "Request compressed response")
}

func RegisterCookieFlag(flags *pflag.FlagSet) {
	flags.StringArrayP("cookie", "b", nil, "Send cookies from string/file")
}

func RegisterCookieJarFlag(flags *pflag.FlagSet) {
	flags.StringP("cookie-jar", "c", "", "Write cookies to <file> after operation")
}

func RegisterEtagCompareFlag(flags *pflag.FlagSet) {
	flags.String("etag-compare", "", "Pass an ETag to compare on conditional requests")
}

func RegisterEtagSaveFlag(flags *pflag.FlagSet) {
	flags.String("etag-save", "", "Parse and save an ETag in the HTTP response")
}

func RegisterHaproxyProtocolFlag(flags *pflag.FlagSet) {
	flags.Bool("haproxy-protocol", false, "Send HAProxy PROXY protocol v1 header")
}

func RegisterHeaderFlag(flags *pflag.FlagSet) {
	flags.StringArrayP("header", "H", nil, "Pass custom header(s) to server")
}

func RegisterJunkSessionCookiesFlag(flags *pflag.FlagSet) {
	flags.BoolP("junk-session-cookies", "j", false, "Ignore session cookies read from file")
}

func RegisterRefererFlag(flags *pflag.FlagSet) {
	flags.StringP("referer", "e", "", "Send Referer Page information")
}

func RegisterTimeCondFlag(flags *pflag.FlagSet) {
	flags.StringP("time-cond", "z", "", "Transfer based on a time condition")
}

func RegisterTrEncodingFlag(flags *pflag.FlagSet) {
	flags.Bool("tr-encoding", false, "Request compressed transfer encoding")
}

func RegisterUserAgentFlag(flags *pflag.FlagSet) {
	flags.StringP("user-agent", "A", "", "Send User-Agent <name> to server")
}

func RegisterAppendFlag(flags *pflag.FlagSet) {
	flags.BoolP("append", "a", false, "Append to target file(s) instead of overwriting")
}

func RegisterContinueAtFlag(flags *pflag.FlagSet) {
	flags.Int64P("continue-at", "C", 0, "Resumed transfer offset")
}

func RegisterCrlfFlag(flags *pflag.FlagSet) {
	flags.Bool("crlf", false, "Convert LF to CRLF in upload")
}

func RegisterDataFlag(flags *pflag.FlagSet) {
	flags.StringP("data", "d", "", "HTTP POST data")
}

func RegisterDataAsciiFlag(flags *pflag.FlagSet) {
	flags.String("data-ascii", "", "HTTP POST ASCII data")
}

func RegisterDataBinaryFlag(flags *pflag.FlagSet) {
	flags.String("data-binary", "", "HTTP POST binary data")
}

func RegisterDataRawFlag(flags *pflag.FlagSet) {
	flags.String("data-raw", "", "HTTP POST data, @ is not special")
}

func RegisterDataUrlencodeFlag(flags *pflag.FlagSet) {
	flags.String("data-urlencode", "", "HTTP POST data URL encoded")
}

func RegisterFormFlag(flags *pflag.FlagSet) {
	flags.StringToStringP("form", "F", map[string]string{}, "Specify multipart MIME data")
}

func RegisterFormEscapeFlag(flags *pflag.FlagSet) {
	flags.Bool("form-escape", false, "Escape multipart form field/file names")
}

func RegisterFormStringFlag(flags *pflag.FlagSet) {
	flags.String("form-string", "", "Specify multipart MIME data")
}

func RegisterJsonFlag(flags *pflag.FlagSet) {
	flags.String("json", "", "HTTP POST JSON")
}

func RegisterUploadFileFlag(flags *pflag.FlagSet) {
	flags.StringP("upload-file", "T", "", "Transfer local FILE to destination")
}

func RegisterDumpHeaderFlag(flags *pflag.FlagSet) {
	flags.StringP("dump-header", "D", "", "Write the received headers to <file>")
}

func RegisterFailFlag(flags *pflag.FlagSet) {
	flags.BoolP("fail", "f", false, "Fail fast with no output on HTTP errors")
}

func RegisterFailEarlyFlag(flags *pflag.FlagSet) {
	flags.Bool("fail-early", false, "Fail on first transfer error, do not continue")
}

func RegisterFailWithBodyFlag(flags *pflag.FlagSet) {
	flags.Bool("fail-with-body", false, "Fail on HTTP errors but save the body")
}

func RegisterNoBufferFlag(flags *pflag.FlagSet) {
	flags.BoolP("no-buffer", "N", false, "Disable buffering of the output stream")
}

func RegisterNoProgressMeterFlag(flags *pflag.FlagSet) {
	flags.Bool("no-progress-meter", false, "Do not show the progress meter")
}

func RegisterOutputFlag(flags *pflag.FlagSet) {
	flags.StringP("output", "o", "", "Write to file instead of stdout")
}

func RegisterOutputDirFlag(flags *pflag.FlagSet) {
	flags.String("output-dir", "", "Directory to save downloads")
}

func RegisterProgressBarFlag(flags *pflag.FlagSet) {
	flags.BoolP("progress-bar", "#", false, "Display transfer progress as a bar")
}

func RegisterShowErrorFlag(flags *pflag.FlagSet) {
	flags.BoolP("show-error", "S", false, "Show error even when -s is used")
}

func RegisterShowHeadersFlag(flags *pflag.FlagSet) {
	flags.BoolP("show-headers", "i", false, "Include protocol response headers in the output")
}

func RegisterSilentFlag(flags *pflag.FlagSet) {
	flags.BoolP("silent", "s", false, "Silent mode")
}

func RegisterStderrFlag(flags *pflag.FlagSet) {
	flags.String("stderr", "", "Where to redirect stderr")
}

func RegisterStyledOutputFlag(flags *pflag.FlagSet) {
	flags.Bool("styled-output", false, "Enable styled output for HTTP headers")
}

func RegisterVerboseFlag(flags *pflag.FlagSet) {
	flags.BoolP("verbose", "v", false, "Make the operation more talkative")
}

func RegisterWriteOutFlag(flags *pflag.FlagSet) {
	flags.StringP("write-out", "w", "", "Use output FORMAT after completion")
}

func RegisterLimitRateFlag(flags *pflag.FlagSet) {
	flags.String("limit-rate", "", "Limit transfer speed to <speed>")
}

func RegisterMaxFilesizeFlag(flags *pflag.FlagSet) {
	flags.Int64("max-filesize", 0, "Maximum file size to download")
}

func RegisterParallelFlag(flags *pflag.FlagSet) {
	flags.BoolP("parallel", "Z", false, "Perform transfers in parallel")
}

func RegisterParallelImmediateFlag(flags *pflag.FlagSet) {
	flags.Bool("parallel-immediate", false, "Do not wait for multiplexing")
}

func RegisterParallelMaxFlag(flags *pflag.FlagSet) {
	flags.Int("parallel-max", 0, "Maximum concurrent parallel transfers")
}

func RegisterRangeFlag(flags *pflag.FlagSet) {
	flags.StringP("range", "r", "", "Retrieve only the bytes within RANGE")
}

func RegisterRateFlag(flags *pflag.FlagSet) {
	flags.String("rate", "", "Request rate for serial transfers")
}

func RegisterRetryFlag(flags *pflag.FlagSet) {
	flags.Int("retry", 0, "Retry request if transient problems occur")
}

func RegisterRetryAllErrorsFlag(flags *pflag.FlagSet) {
	flags.Bool("retry-all-errors", false, "Retry on all errors (use with --retry)")
}

func RegisterRetryConnrefusedFlag(flags *pflag.FlagSet) {
	flags.Bool("retry-connrefused", false, "Retry on connection refused (use with --retry)")
}

func RegisterRetryDelayFlag(flags *pflag.FlagSet) {
	flags.Int("retry-delay", 0, "Wait time between retries")
}

func RegisterRetryMaxTimeFlag(flags *pflag.FlagSet) {
	flags.Int("retry-max-time", 0, "Retry only within this period")
}

func RegisterSpeedLimitFlag(flags *pflag.FlagSet) {
	flags.IntP("speed-limit", "Y", 0, "Stop transfers slower than this")
}

func RegisterSpeedTimeFlag(flags *pflag.FlagSet) {
	flags.IntP("speed-time", "y", 0, "Trigger 'speed-limit' abort after this time")
}

func RegisterCreateDirsFlag(flags *pflag.FlagSet) {
	flags.Bool("create-dirs", false, "Create necessary local directory hierarchy")
}

func RegisterCreateFileModeFlag(flags *pflag.FlagSet) {
	flags.String("create-file-mode", "", "File mode for created files")
}

func RegisterNoClobberFlag(flags *pflag.FlagSet) {
	flags.Bool("no-clobber", false, "Do not overwrite files")
}

func RegisterRemoteHeaderNameFlag(flags *pflag.FlagSet) {
	flags.BoolP("remote-header-name", "J", false, "Use the header-provided filename")
}

func RegisterRemoteNameFlag(flags *pflag.FlagSet) {
	flags.BoolP("remote-name", "O", false, "Write output to a file named as the remote file")
}

func RegisterRemoteTimeFlag(flags *pflag.FlagSet) {
	flags.BoolP("remote-time", "R", false, "Set remote file's time on local output")
}

func RegisterRemoveOnErrorFlag(flags *pflag.FlagSet) {
	flags.Bool("remove-on-error", false, "Remove output file on errors")
}

func RegisterSkipExistingFlag(flags *pflag.FlagSet) {
	flags.Bool("skip-existing", false, "Skip download if local file already exists")
}

func RegisterXattrFlag(flags *pflag.FlagSet) {
	flags.Bool("xattr", false, "Store metadata in extended file attributes")
}

func RegisterDisableEprtFlag(flags *pflag.FlagSet) {
	flags.Bool("disable-eprt", false, "Inhibit using EPRT or LPRT")
}

func RegisterDisableEpsvFlag(flags *pflag.FlagSet) {
	flags.Bool("disable-epsv", false, "Inhibit using EPSV")
}

func RegisterFtpAccountFlag(flags *pflag.FlagSet) {
	flags.String("ftp-account", "", "Account data string")
}

func RegisterFtpAlternativeToUserFlag(flags *pflag.FlagSet) {
	flags.String("ftp-alternative-to-user", "", "String to replace USER [name]")
}

func RegisterFtpCreateDirsFlag(flags *pflag.FlagSet) {
	flags.Bool("ftp-create-dirs", false, "Create the remote dirs if not present")
}

func RegisterFtpMethodFlag(flags *pflag.FlagSet) {
	flags.String("ftp-method", "", "Specify CWD method")
}

func RegisterFtpPasvFlag(flags *pflag.FlagSet) {
	flags.Bool("ftp-pasv", false, "Use PASV/EPSV instead of PORT")
}

func RegisterFtpPortFlag(flags *pflag.FlagSet) {
	flags.StringP("ftp-port", "P", "", "Use PORT instead of PASV")
}

func RegisterFtpPretFlag(flags *pflag.FlagSet) {
	flags.Bool("ftp-pret", false, "Send PRET before PASV")
}

func RegisterFtpSkipPasvIpFlag(flags *pflag.FlagSet) {
	flags.Bool("ftp-skip-pasv-ip", false, "Skip the IP address for PASV")
}

func RegisterFtpSslCccFlag(flags *pflag.FlagSet) {
	flags.Bool("ftp-ssl-ccc", false, "Send CCC after authenticating")
}

func RegisterFtpSslCccModeFlag(flags *pflag.FlagSet) {
	flags.String("ftp-ssl-ccc-mode", "", "Set CCC mode")
}

func RegisterFtpSslControlFlag(flags *pflag.FlagSet) {
	flags.Bool("ftp-ssl-control", false, "Require SSL/TLS for FTP login, clear for transfer")
}

func RegisterListOnlyFlag(flags *pflag.FlagSet) {
	flags.BoolP("list-only", "l", false, "List only mode")
}

func RegisterPreproxyFlag(flags *pflag.FlagSet) {
	flags.String("preproxy", "", "Go via the proxy first")
}

func RegisterProtoFlag(flags *pflag.FlagSet) {
	flags.String("proto", "", "Enable/disable PROTOCOLS")
}

func RegisterProtoDefaultFlag(flags *pflag.FlagSet) {
	flags.String("proto-default", "", "Use PROTOCOL for any URL missing a scheme")
}

func RegisterProtoRedirFlag(flags *pflag.FlagSet) {
	flags.String("proto-redir", "", "Enable/disable PROTOCOLS on redirect")
}

func RegisterSocks4Flag(flags *pflag.FlagSet) {
	flags.String("socks4", "", "SOCKS4 proxy on given host + port")
}

func RegisterSocks4aFlag(flags *pflag.FlagSet) {
	flags.String("socks4a", "", "SOCKS4a proxy on given host + port")
}

func RegisterSocks5Flag(flags *pflag.FlagSet) {
	flags.String("socks5", "", "SOCKS5 proxy on given host + port")
}

func RegisterSocks5GssapiNecFlag(flags *pflag.FlagSet) {
	flags.Bool("socks5-gssapi-nec", false, "Compatibility with NEC SOCKS5 server")
}

func RegisterSocks5GssapiServiceFlag(flags *pflag.FlagSet) {
	flags.String("socks5-gssapi-service", "", "SOCKS5 proxy service name for GSS-API")
}

func RegisterSocks5HostnameFlag(flags *pflag.FlagSet) {
	flags.String("socks5-hostname", "", "SOCKS5 proxy, pass host name to proxy")
}

func RegisterSslFlag(flags *pflag.FlagSet) {
	flags.Bool("ssl", false, "Try SSL/TLS")
}

func RegisterSslAutoClientCertFlag(flags *pflag.FlagSet) {
	flags.Bool("ssl-auto-client-cert", false, "Try to find and use auto client certificate")
}

func RegisterSslNoRevokeFlag(flags *pflag.FlagSet) {
	flags.Bool("ssl-no-revoke", false, "Disable cert revocation checks (WinSSL)")
}

func RegisterSslReqdFlag(flags *pflag.FlagSet) {
	flags.Bool("ssl-reqd", false, "Require SSL/TLS")
}

func RegisterSslRevokeBestEffortFlag(flags *pflag.FlagSet) {
	flags.Bool("ssl-revoke-best-effort", false, "Ignore missing/offline cert distribution points")
}

func RegisterSslv2Flag(flags *pflag.FlagSet) {
	flags.BoolP("sslv2", "2", false, "Use SSLv2 (deprecated)")
}

func RegisterSslv3Flag(flags *pflag.FlagSet) {
	flags.BoolP("sslv3", "3", false, "Use SSLv3 (deprecated)")
}

func RegisterTraceFlag(flags *pflag.FlagSet) {
	flags.String("trace", "", "Write a debug trace to FILE")
}

func RegisterTraceAsciiFlag(flags *pflag.FlagSet) {
	flags.String("trace-ascii", "", "Like --trace, but without hex output")
}

func RegisterTraceConfigFlag(flags *pflag.FlagSet) {
	flags.String("trace-config", "", "Details to log in trace/verbose")
}

func RegisterTraceIdsFlag(flags *pflag.FlagSet) {
	flags.Bool("trace-ids", false, "Add transfer IDs to trace/verbose output")
}

func RegisterTraceTimeFlag(flags *pflag.FlagSet) {
	flags.Bool("trace-time", false, "Add time stamps to trace/verbose output")
}

func RegisterAbstractUnixSocketFlag(flags *pflag.FlagSet) {
	flags.String("abstract-unix-socket", "", "Connect via abstract Unix domain socket")
}

func RegisterAltSvcFlag(flags *pflag.FlagSet) {
	flags.String("alt-svc", "", "Enable alt-svc with this cache file")
}

func RegisterConfigFlag(flags *pflag.FlagSet) {
	flags.StringP("config", "K", "", "Read config from a file")
}

func RegisterCrlfileFlag(flags *pflag.FlagSet) {
	flags.String("crlfile", "", "Use a CRL list")
}

func RegisterDisableFlag(flags *pflag.FlagSet) {
	flags.BoolP("disable", "q", false, "Disable .curlrc")
}

func RegisterDohCertStatusFlag(flags *pflag.FlagSet) {
	flags.Bool("doh-cert-status", false, "Verify DoH server certificate status")
}

func RegisterEchFlag(flags *pflag.FlagSet) {
	flags.String("ech", "", "Configure ECH")
}

func RegisterEgdFileFlag(flags *pflag.FlagSet) {
	flags.String("egd-file", "", "EGD socket path for entropy")
}

func RegisterEngineFlag(flags *pflag.FlagSet) {
	flags.String("engine", "", "OpenSSL crypto engine to use")
}

func RegisterHostpubmd5Flag(flags *pflag.FlagSet) {
	flags.String("hostpubmd5", "", "Acceptable MD5 hash of the host public key")
}

func RegisterHstsFlag(flags *pflag.FlagSet) {
	flags.String("hsts", "", "Enable HSTS with this cache file")
}

func RegisterIgnoreContentLengthFlag(flags *pflag.FlagSet) {
	flags.Bool("ignore-content-length", false, "Ignore the size of the remote resource")
}

func RegisterLoginOptionsFlag(flags *pflag.FlagSet) {
	flags.String("login-options", "", "Server login options")
}

func RegisterMailAuthFlag(flags *pflag.FlagSet) {
	flags.String("mail-auth", "", "Originator address of the original email")
}

func RegisterMailFromFlag(flags *pflag.FlagSet) {
	flags.String("mail-from", "", "Mail from this address")
}

func RegisterMailRcptFlag(flags *pflag.FlagSet) {
	flags.String("mail-rcpt", "", "Mail to this address")
}

func RegisterMetalinkFlag(flags *pflag.FlagSet) {
	flags.Bool("metalink", false, "Process given URLs as metalink XML file")
}

func RegisterNextFlag(flags *pflag.FlagSet) {
	flags.Bool("next", false, "Allows several URLs in a single command line")
}

func RegisterNoNpnFlag(flags *pflag.FlagSet) {
	flags.Bool("no-npn", false, "Disable the NPN TLS extension")
}

func RegisterPubkeyFlag(flags *pflag.FlagSet) {
	flags.String("pubkey", "", "Public key file name")
}

func RegisterRandomFileFlag(flags *pflag.FlagSet) {
	flags.String("random-file", "", "File for reading random data from")
}

func RegisterSuppressConnectHeadersFlag(flags *pflag.FlagSet) {
	flags.Bool("suppress-connect-headers", false, "Suppress proxy CONNECT response headers")
}

func RegisterTelnetOptionFlag(flags *pflag.FlagSet) {
	flags.String("telnet-option", "", "Set telnet option")
}

func RegisterTftpNoOptionsFlag(flags *pflag.FlagSet) {
	flags.Bool("tftp-no-options", false, "Do not send any TFTP options")
}

func RegisterUnixSocketFlag(flags *pflag.FlagSet) {
	flags.String("unix-socket", "", "Connect through this Unix domain socket")
}

func RegisterUseAsciiFlag(flags *pflag.FlagSet) {
	flags.BoolP("use-ascii", "B", false, "Use ASCII/text transfer")
}

func RegisterVariableFlag(flags *pflag.FlagSet) {
	flags.String("variable", "", "Set variable")
}

func RegisterTimeoutFlags(flags *pflag.FlagSet) {
	RegisterConnectTimeoutFlag(flags)
	RegisterMaxTimeFlag(flags)
}

func RegisterProxyFlags(flags *pflag.FlagSet) {
	RegisterProxyFlag(flags)
	RegisterNoproxyFlag(flags)
	RegisterProxyAnyauthFlag(flags)
	RegisterProxyBasicFlag(flags)
	RegisterProxyCacertFlag(flags)
	RegisterProxyCapathFlag(flags)
	RegisterProxyCertFlag(flags)
	RegisterProxyCertTypeFlag(flags)
	RegisterProxyCiphersFlag(flags)
	RegisterProxyCrlfileFlag(flags)
	RegisterProxyDigestFlag(flags)
	RegisterProxyHeaderFlag(flags)
	RegisterProxyHttp2Flag(flags)
	RegisterProxyInsecureFlag(flags)
	RegisterProxyKeyFlag(flags)
	RegisterProxyKeyTypeFlag(flags)
	RegisterProxyNegotiateFlag(flags)
	RegisterProxyNtlmFlag(flags)
	RegisterProxyPassFlag(flags)
	RegisterProxyServiceNameFlag(flags)
	RegisterProxySslAllowBeastFlag(flags)
	RegisterProxyTls13CiphersFlag(flags)
	RegisterProxyTlsauthtypeFlag(flags)
	RegisterProxyTlspasswordFlag(flags)
	RegisterProxyTlsuserFlag(flags)
	RegisterProxyTlsv1Flag(flags)
	RegisterProxyUserFlag(flags)
	RegisterProxy10Flag(flags)
	RegisterProxytunnelFlag(flags)
}

func RegisterTLSFlags(flags *pflag.FlagSet) {
	RegisterInsecureFlag(flags)
	RegisterCacertFlag(flags)
	RegisterCapathFlag(flags)
	RegisterCaNativeFlag(flags)
	RegisterCertFlag(flags)
	RegisterCertStatusFlag(flags)
	RegisterCertTypeFlag(flags)
	RegisterCiphersFlag(flags)
	RegisterCurvesFlag(flags)
	RegisterFalseStartFlag(flags)
	RegisterKeyFlag(flags)
	RegisterKeyTypeFlag(flags)
	RegisterNoAlpnFlag(flags)
	RegisterNoSessionidFlag(flags)
	RegisterPassFlag(flags)
	RegisterPinnedpubkeyFlag(flags)
	RegisterSslAllowBeastFlag(flags)
	RegisterTlsMaxFlag(flags)
	RegisterTls13CiphersFlag(flags)
	RegisterTlsauthtypeFlag(flags)
	RegisterTlspasswordFlag(flags)
	RegisterTlsuserFlag(flags)
	RegisterTlsv1Flag(flags)
	RegisterTlsv10Flag(flags)
	RegisterTlsv11Flag(flags)
	RegisterTlsv12Flag(flags)
	RegisterTlsv13Flag(flags)
}

func RegisterRedirectFlags(flags *pflag.FlagSet) {
	RegisterLocationFlag(flags)
	RegisterLocationTrustedFlag(flags)
	RegisterMaxRedirsFlag(flags)
	RegisterPost301Flag(flags)
	RegisterPost302Flag(flags)
	RegisterPost303Flag(flags)
}

func RegisterNetworkFlags(flags *pflag.FlagSet) {
	RegisterConnectToFlag(flags)
	RegisterDohInsecureFlag(flags)
	RegisterDohUrlFlag(flags)
	RegisterHappyEyeballsTimeoutMsFlag(flags)
	RegisterInterfaceFlag(flags)
	RegisterIpv4Flag(flags)
	RegisterIpv6Flag(flags)
	RegisterLocalPortFlag(flags)
	RegisterResolveFlag(flags)
}

func RegisterKeepAliveFlags(flags *pflag.FlagSet) {
	RegisterExpect100TimeoutFlag(flags)
	RegisterKeepaliveTimeFlag(flags)
	RegisterNoKeepaliveFlag(flags)
	RegisterTcpFastopenFlag(flags)
	RegisterTcpNodelayFlag(flags)
}

func RegisterAuthFlags(flags *pflag.FlagSet) {
	RegisterAnyauthFlag(flags)
	RegisterAwsSigv4Flag(flags)
	RegisterBasicFlag(flags)
	RegisterDelegationFlag(flags)
	RegisterDigestFlag(flags)
	RegisterKrbFlag(flags)
	RegisterNegotiateFlag(flags)
	RegisterNetrcFlag(flags)
	RegisterNetrcFileFlag(flags)
	RegisterNetrcOptionalFlag(flags)
	RegisterNtlmFlag(flags)
	RegisterNtlmWbFlag(flags)
	RegisterOauth2BearerFlag(flags)
	RegisterSaslIrFlag(flags)
	RegisterServiceNameFlag(flags)
	RegisterUserFlag(flags)
}

func RegisterRequestFlags(flags *pflag.FlagSet) {
	RegisterGetFlag(flags)
	RegisterGloboffFlag(flags)
	RegisterHeadFlag(flags)
	RegisterHttp09Flag(flags)
	RegisterHttp10Flag(flags)
	RegisterHttp11Flag(flags)
	RegisterHttp2Flag(flags)
	RegisterHttp2PriorKnowledgeFlag(flags)
	RegisterHttp3Flag(flags)
	RegisterHttp3OnlyFlag(flags)
	RegisterPathAsIsFlag(flags)
	RegisterRawFlag(flags)
	RegisterRequestFlag(flags)
	RegisterRequestTargetFlag(flags)
	RegisterUrlFlag(flags)
}

func RegisterHeaderFlags(flags *pflag.FlagSet) {
	RegisterCompressedFlag(flags)
	RegisterCookieFlag(flags)
	RegisterCookieJarFlag(flags)
	RegisterEtagCompareFlag(flags)
	RegisterEtagSaveFlag(flags)
	RegisterHaproxyProtocolFlag(flags)
	RegisterHeaderFlag(flags)
	RegisterJunkSessionCookiesFlag(flags)
	RegisterRefererFlag(flags)
	RegisterTimeCondFlag(flags)
	RegisterTrEncodingFlag(flags)
	RegisterUserAgentFlag(flags)
}

func RegisterBodyFlags(flags *pflag.FlagSet) {
	RegisterAppendFlag(flags)
	RegisterContinueAtFlag(flags)
	RegisterCrlfFlag(flags)
	RegisterDataFlag(flags)
	RegisterDataAsciiFlag(flags)
	RegisterDataBinaryFlag(flags)
	RegisterDataRawFlag(flags)
	RegisterDataUrlencodeFlag(flags)
	RegisterFormFlag(flags)
	RegisterFormEscapeFlag(flags)
	RegisterFormStringFlag(flags)
	RegisterJsonFlag(flags)
	RegisterUploadFileFlag(flags)
}

func RegisterOutputFlags(flags *pflag.FlagSet) {
	RegisterDumpHeaderFlag(flags)
	RegisterFailFlag(flags)
	RegisterFailEarlyFlag(flags)
	RegisterFailWithBodyFlag(flags)
	RegisterNoBufferFlag(flags)
	RegisterNoProgressMeterFlag(flags)
	RegisterOutputFlag(flags)
	RegisterOutputDirFlag(flags)
	RegisterProgressBarFlag(flags)
	RegisterShowErrorFlag(flags)
	RegisterShowHeadersFlag(flags)
	RegisterSilentFlag(flags)
	RegisterStderrFlag(flags)
	RegisterStyledOutputFlag(flags)
	RegisterVerboseFlag(flags)
	RegisterWriteOutFlag(flags)
}

func RegisterTransferFlags(flags *pflag.FlagSet) {
	RegisterLimitRateFlag(flags)
	RegisterMaxFilesizeFlag(flags)
	RegisterParallelFlag(flags)
	RegisterParallelImmediateFlag(flags)
	RegisterParallelMaxFlag(flags)
	RegisterRangeFlag(flags)
	RegisterRateFlag(flags)
	RegisterRetryFlag(flags)
	RegisterRetryAllErrorsFlag(flags)
	RegisterRetryConnrefusedFlag(flags)
	RegisterRetryDelayFlag(flags)
	RegisterRetryMaxTimeFlag(flags)
	RegisterSpeedLimitFlag(flags)
	RegisterSpeedTimeFlag(flags)
}

func RegisterFileFlags(flags *pflag.FlagSet) {
	RegisterCreateDirsFlag(flags)
	RegisterCreateFileModeFlag(flags)
	RegisterNoClobberFlag(flags)
	RegisterRemoteHeaderNameFlag(flags)
	RegisterRemoteNameFlag(flags)
	RegisterRemoteTimeFlag(flags)
	RegisterRemoveOnErrorFlag(flags)
	RegisterSkipExistingFlag(flags)
	RegisterXattrFlag(flags)
}

func RegisterFTPFlags(flags *pflag.FlagSet) {
	RegisterDisableEprtFlag(flags)
	RegisterDisableEpsvFlag(flags)
	RegisterFtpAccountFlag(flags)
	RegisterFtpAlternativeToUserFlag(flags)
	RegisterFtpCreateDirsFlag(flags)
	RegisterFtpMethodFlag(flags)
	RegisterFtpPasvFlag(flags)
	RegisterFtpPortFlag(flags)
	RegisterFtpPretFlag(flags)
	RegisterFtpSkipPasvIpFlag(flags)
	RegisterFtpSslCccFlag(flags)
	RegisterFtpSslCccModeFlag(flags)
	RegisterFtpSslControlFlag(flags)
	RegisterListOnlyFlag(flags)
}

func RegisterProtocolFlags(flags *pflag.FlagSet) {
	RegisterPreproxyFlag(flags)
	RegisterProtoFlag(flags)
	RegisterProtoDefaultFlag(flags)
	RegisterProtoRedirFlag(flags)
	RegisterSocks4Flag(flags)
	RegisterSocks4aFlag(flags)
	RegisterSocks5Flag(flags)
	RegisterSocks5GssapiNecFlag(flags)
	RegisterSocks5GssapiServiceFlag(flags)
	RegisterSocks5HostnameFlag(flags)
	RegisterSslFlag(flags)
	RegisterSslAutoClientCertFlag(flags)
	RegisterSslNoRevokeFlag(flags)
	RegisterSslReqdFlag(flags)
	RegisterSslRevokeBestEffortFlag(flags)
	RegisterSslv2Flag(flags)
	RegisterSslv3Flag(flags)
}

func RegisterTraceFlags(flags *pflag.FlagSet) {
	RegisterTraceFlag(flags)
	RegisterTraceAsciiFlag(flags)
	RegisterTraceConfigFlag(flags)
	RegisterTraceIdsFlag(flags)
	RegisterTraceTimeFlag(flags)
}

func RegisterMiscFlags(flags *pflag.FlagSet) {
	RegisterAbstractUnixSocketFlag(flags)
	RegisterAltSvcFlag(flags)
	RegisterConfigFlag(flags)
	RegisterCrlfileFlag(flags)
	RegisterDisableFlag(flags)
	RegisterDohCertStatusFlag(flags)
	RegisterEchFlag(flags)
	RegisterEgdFileFlag(flags)
	RegisterEngineFlag(flags)
	RegisterHostpubmd5Flag(flags)
	RegisterHstsFlag(flags)
	RegisterIgnoreContentLengthFlag(flags)
	RegisterLoginOptionsFlag(flags)
	RegisterMailAuthFlag(flags)
	RegisterMailFromFlag(flags)
	RegisterMailRcptFlag(flags)
	RegisterMetalinkFlag(flags)
	RegisterNextFlag(flags)
	RegisterNoNpnFlag(flags)
	RegisterPubkeyFlag(flags)
	RegisterRandomFileFlag(flags)
	RegisterSuppressConnectHeadersFlag(flags)
	RegisterTelnetOptionFlag(flags)
	RegisterTftpNoOptionsFlag(flags)
	RegisterUnixSocketFlag(flags)
	RegisterUseAsciiFlag(flags)
	RegisterVariableFlag(flags)
}

func RegisterClientFlags(flags *pflag.FlagSet) {
	RegisterTimeoutFlags(flags)
	RegisterProxyFlags(flags)
	RegisterTLSFlags(flags)
	RegisterRedirectFlags(flags)
	RegisterNetworkFlags(flags)
	RegisterKeepAliveFlags(flags)
}

func RegisterFlags(flags *pflag.FlagSet) {
	RegisterClientFlags(flags)
	RegisterAuthFlags(flags)
	RegisterRequestFlags(flags)
	RegisterHeaderFlags(flags)
	RegisterBodyFlags(flags)
	RegisterOutputFlags(flags)
	RegisterTransferFlags(flags)
	RegisterFileFlags(flags)
	RegisterFTPFlags(flags)
	RegisterProtocolFlags(flags)
	RegisterTraceFlags(flags)
	RegisterMiscFlags(flags)
}
