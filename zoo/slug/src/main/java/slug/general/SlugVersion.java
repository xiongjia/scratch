package slug.general;

public class SlugVersion {
  private int major;
  private int minor;

  public int getMajor() {
    return major;
  }

  public void setMajor(int major) {
    this.major = major;
  }

  public int getMinor() {
    return minor;
  }

  public void setMinor(int minor) {
    this.minor = minor;
  }

  public String getFullVersion() {
    return String.format("%d.%d", getMajor(), getMinor());
  }

  @Override
  public String toString() {
    return "SlugVersion{"
      + "major=" + major
      + ", minor=" + minor
      + '}';
  }
}
