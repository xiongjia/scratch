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

const char* alog_getfilebasename(const char *filename) {
    const char *slash = strrchr(filename, '/');
    if (slash) {
        return slash + 1;
    }
    else {
        return filename;
    }
}

#pragma mark - Logger Interfaces

+ (id)logger {
    /* TODO create a logger level for this class */
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
    if (!format) {
        return;
    }
    
    va_list args;
    va_start(args, format);
    NSString *logMsg = [[NSString alloc] initWithFormat:format arguments:args];
    va_end(args);
    
    AbathurLogger *inst = [AbathurLogger logger];
    [inst logMessage: flag
             srcFile: srcFile
             srcLine: srcLine
             message: logMsg];
}

#pragma mark - Logger functions

- (void)logMessage: (AbathurLogFlag)flag
          srcFile: (const char *)srcFile
          srcLine: (int)srcLine
          message: (NSString*)message {
    const char *srcBaseFilename = alog_getfilebasename(srcFile);
    NSLog(@"%s:%d - %@", srcBaseFilename, srcLine, message);
}

@end
