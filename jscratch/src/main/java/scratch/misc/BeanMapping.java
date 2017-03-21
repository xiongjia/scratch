package scratch.misc;

import org.modelmapper.ModelMapper;
import org.modelmapper.PropertyMap;

public class BeanMapping  {
  private static final PropertyMap<BeanSource, BeanTarget> srcToTarget =
      new PropertyMap<BeanSource, BeanTarget>() {
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

  public static BeanTarget toTarget(BeanSource src) {
    return modelMapper.map(src, BeanTarget.class);
  }
}
