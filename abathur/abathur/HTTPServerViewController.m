/* HTTPServerViewController.m - Abathur */

#import "HTTPServerViewController.h"
#import "AbathurLogger.h"
#import "HTTPServer.h"

@interface HTTPServerViewController ()

@end

@implementation HTTPServerViewController

- (void)viewDidLoad {
    [super viewDidLoad];
    ABLOG(ABLOG_DBG, @"HTTPServerView didLoad");
    HTTPServer *httpServ = [HTTPServer httpServer];
    NSString *ip4Addr = [httpServ getIPAddress: YES];
    ABLOG(ABLOG_DBG, @"Local Addr %@", ip4Addr);
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
}

@end
