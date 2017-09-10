/**
 * OpenSSL tests
 */

#include <stdio.h>
#include "openssl/ssl.h"
#include "openssl/bio.h"
#include "openssl/err.h"

int main(const int argc, const char **argv) {
    // ERR_load_BIO_strings();
    SSL_load_error_strings();
    OpenSSL_add_all_algorithms();

    SSL_CTX *sslCtx = SSL_CTX_new(SSLv23_client_method());
    int ret = SSL_CTX_load_verify_locations(sslCtx, "TrustStore.pem", NULL);
    if (0 == ret) {
        ERR_print_errors_fp(stderr);
    }

    SSL_CTX_free(sslCtx);
    sslCtx = nullptr;
    return 0;
}
