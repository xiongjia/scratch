//
//  AbathurLogger.h
//  abathur
//
//  Created by LeXiongJia on 4/12/15.
//  Copyright (c) 2015 LeXiongJia. All rights reserved.
//

#import <Foundation/Foundation.h>


#pragma mark - Logger Macro
typedef int AbathurLogFlag;
static const AbathurLogFlag ABLOG_ERR = 1 << 0;
static const AbathurLogFlag ABLOG_WAN = 1 << 1;
static const AbathurLogFlag ABLOG_INF = 1 << 2;
static const AbathurLogFlag ABLOG_DBG = 1 << 3;

#define ABLOG(_flag, _format, ...) \
    [AbathurLogger log: _flag \
               srcFile: __FILE__ \
               srcLine: __LINE__ \
                format: (_format), ##__VA_ARGS__]

#pragma mark - Logger class

@interface AbathurLogger : NSObject

+ (id)logger;

+ (void)log: (AbathurLogFlag)flag
    srcFile: (const char *)srcFile
    srcLine: (int)srcLine
     format: (NSString *)format, ...;

@end
