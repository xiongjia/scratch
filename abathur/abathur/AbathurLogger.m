//
//  AbathurLogger.m
//  abathur
//
//  Created by LeXiongJia on 4/12/15.
//  Copyright (c) 2015 LeXiongJia. All rights reserved.
//

#import "AbathurLogger.h"

@implementation AbathurLogger

#pragma mark - Misc



#pragma mark - Logger Interfaces

+ (id)logger {
    static AbathurLogger *inst = nil;
    static dispatch_once_t onceToken;
    dispatch_once(&onceToken, ^{
        inst = [[self alloc] init];
    });
    return inst;
}

+ (void)log: (AbathurLogFlag)flag
    srcFile: (const char *)srcFile
    srcLine: (int)srcLine
     format: (NSString *)format, ... {
    // AbathurLogger *inst = [AbathurLogger logger];
    
    va_list args;
    if (format) {
        va_start(args, format);
        NSString *logMsg = [[NSString alloc] initWithFormat:format arguments:args];
        va_end(args);
        NSLog(@"%@", logMsg);
    }
}

@end
