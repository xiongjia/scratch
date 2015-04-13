/* HTTPServer.m - Abathur */

#import "HTTPServer.h"
#import "AbathurLogger.h"

@implementation HTTPServer

#pragma mark - The constant values

// static const unsigned int HTTP_SERV_DEFAULT_PORT = 8080;

#pragma mark - Interfaces

+ (id)httpServer {
    static HTTPServer *inst = nil;
    static dispatch_once_t onceToken;
    dispatch_once(&onceToken, ^{
        inst = [[self alloc] init];
    });
    return inst;
}

@end
