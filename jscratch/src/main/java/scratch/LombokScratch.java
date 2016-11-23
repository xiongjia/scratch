package scratch;

import lombok.Data;
import lombok.Getter;
import lombok.Setter;

/**
 * Scratch code for the Lombok https://github.com/rzwitserloot/lombok
 * NOTE:
 *   In Java IDE (e.g. Eclipse), you need install the Lombok first. 
 *   (Run java -jar lombok.jar)
 *   
 * @author lexiongjia@gmail.com
 */
public class LombokScratch {
    @Data
    public static class LombokData {
    	private String txt = this.getClass().getName();
    	private int num = 0;
    }
 
    public static class LombokSetterGetter {
    	@Setter @Getter private String txt = this.getClass().getName();
    	@Setter @Getter private int num = 0;
    }
};