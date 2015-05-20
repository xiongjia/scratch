--

local logger = require "logger"

-- Logger module
local function main()
    logger.debug('debug log')
    logger.info('info log')
    logger.warn('warn log')
    logger.err('error log')
end
main()

