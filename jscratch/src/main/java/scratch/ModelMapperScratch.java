package scratch;

import org.modelmapper.ModelMapper;
import org.modelmapper.PropertyMap;
import lombok.Data;

/**
 * Scratch code for the ModelMapper https://github.com/jhalterman/modelmapper
 *
 * NOTE:
 *   This sample is using lombok.builder.
 *   In Java IDE (e.g. Eclipse), you need install the Lombok first.
 *   (Run java -jar lombok.jar)
 *
 * @author lexiongjia@gmail.com
 */
public class ModelMapperScratch {
    private static final PropertyMap<Source, Target> srcToTarget = new PropertyMap<Source, Target>() {
        @Override
        protected void configure() {
            map().setTargetStr(source.getSrcStr());
            map().setTargetNum(source.getSrcNum());
        }
    };
    private static final ModelMapper modelMapper = new ModelMapper();

    static {
        modelMapper.addMappings(srcToTarget);
    }

    public static Target toTarget(Source src) {
        return modelMapper.map(src, Target.class);
    }

    @Data
    public static class Source {
        private String srcStr = "source";
        private int srcNum = 0;
    }

    @Data
    public static class Target {
        private String targetStr = "target";
        private int targetNum = 0;
    }
}
