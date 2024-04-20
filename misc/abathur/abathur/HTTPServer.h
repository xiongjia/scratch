/* HTTPServer.h - Abathur */

#import <Foundation/Foundation.h>

@interface HTTPServer : NSObject

+ (id)httpServer;

- (NSString *)getIPAddress:(BOOL)preferIPv4;

@end
