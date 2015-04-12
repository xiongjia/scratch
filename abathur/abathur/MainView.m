/* MainView.m - Abathur */

#import "MainView.h"
#import "AbathurLogger.h"

@interface MainView ()
@property (nonatomic, strong) NSArray *sampleItems;
@end

#pragma mark - The constant values

static NSString *const KEY_TITLE = @"Title";
static NSString *const KEY_DESC  = @"Desc";
static NSString *const KEY_CLASS = @"Class";
static NSString *const SAMPLE_CELL = @"SampleCell";

@implementation MainView

#pragma mark - General functions

- (void)viewDidLoad {
    [super viewDidLoad];
    ABLOG(ABLOG_DBG|ABLOG_INF, @"Main view did load");

    /* All samples */
    self.sampleItems = @[
                         @{ KEY_TITLE: @"HTTP Server",
                            KEY_DESC: @"A simple HTTP sample Server",
                            KEY_CLASS: @"HTTPServerViewController",
                            },
                         ];

}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
}

#pragma mark - Table view data

- (NSInteger)tableView:(UITableView *)tableView
 numberOfRowsInSection:(NSInteger)section {
    return [self.sampleItems count];
}

- (UITableViewCell *)tableView:(UITableView *)tableView
         cellForRowAtIndexPath:(NSIndexPath *)indexPath {
    UITableViewCell *cell = [tableView dequeueReusableCellWithIdentifier: SAMPLE_CELL];
    NSDictionary *sampleInfo = self.sampleItems[indexPath.row];
    cell.textLabel.text = sampleInfo[KEY_TITLE];
    cell.detailTextLabel.text = sampleInfo[KEY_DESC];
    return cell;
}

- (void)tableView:(UITableView *)tableView
    didSelectRowAtIndexPath:(NSIndexPath *)indexPath {
    NSDictionary *sampleInfo = self.sampleItems[indexPath.row];
    Class sampleClass = NSClassFromString(sampleInfo[KEY_CLASS]);
    if (sampleClass) {
        id instance = [[sampleClass alloc] init];
        if ([instance isKindOfClass:[UIViewController class]]) {
            [(UIViewController *)instance setTitle: sampleInfo[KEY_TITLE]];
            [self.navigationController pushViewController: (UIViewController *)instance
                                                 animated: YES];
        }
    }
    [tableView deselectRowAtIndexPath: indexPath
                             animated: YES];
}

@end
