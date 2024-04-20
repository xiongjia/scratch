/**
 */

#ifndef _AB_MOD_LOGGER_H_
#define _AB_MOD_LOGGER_H_ 1

#include "ab_module.h"

_ANUBARAK_BEGIN

class ModLogger : public Module{public:
    ModLogger(void);
    virtual const char* GetModuleName(void) const;

protected:
    virtual const luaL_Reg* GetLUAReg(void);
};

_ANUBARAK_END

#endif /* !defined(_AB_MOD_LOGGER_H_) */
