package xiongjia.kharazim;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class LoggingManager {
    private static final String PACKAGE_PREFIX = "xiongjia.kharazim";

    private LoggingManager() {
        /* non-instantiable - static methods only */
    }

    private static String removePackagePrefix(String name) {
        if (name.startsWith(PACKAGE_PREFIX)) {
            name = name.substring(PACKAGE_PREFIX.length());
        }
        return name;
    }

    public static Log getLoggerForClass() {
        String className = new Exception().getStackTrace()[1].getClassName();
        return getLoggerFor(removePackagePrefix(className));
    }

    public static Log getLoggerFor(String category) {
        return LogFactory.getLog(category);
    }
}
