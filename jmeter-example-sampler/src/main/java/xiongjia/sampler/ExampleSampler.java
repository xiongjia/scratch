package xiongjia.sampler;

import org.apache.jmeter.samplers.AbstractSampler;
import org.apache.jmeter.samplers.Entry;
import org.apache.jmeter.samplers.SampleResult;
import org.apache.jorphan.logging.LoggingManager;
import org.apache.log.Logger;

public class ExampleSampler extends AbstractSampler {
    private static final long serialVersionUID = -4301720123844098677L;

    private static final Logger LOG = LoggingManager.getLoggerForClass();    
    private static final boolean isDbgEnabled = LOG.isDebugEnabled();

    /* sampler prop names */
    public static final String REP_DATA = "XJ.ExampleSampler.RepData";
    
    public ExampleSampler() {
        if (isDbgEnabled) {
            LOG.debug("Creating ExampleSampler");
        }
    }

    /**
     * Do a sampling and return its results.
     *
     * @param entry
     *            <code>Entry</code> to be sampled
     * @return results of the sampling
     */
    @Override
    public SampleResult sample(Entry entry) {
        if (isDbgEnabled) {
            LOG.debug("Creating SampleResult");
        }

        SampleResult res = new SampleResult();
        res.setSampleLabel(this.getName());

        /* Perform the sampling */
        boolean isOK = false;
        String rep = "";
        res.sampleStart();
        try {
            /* TODO: Do tests here ... */
            rep = getPropertyAsString(REP_DATA);
            if (isDbgEnabled) {
                LOG.debug("ExampleSampler REP_DATA = " + rep);
            }

            /* Set up the sample result details */
            res.setSamplerData("Example Data");
            res.setResponseData(rep, null);
            res.setDataType(SampleResult.TEXT);
            res.setResponseCodeOK();
            res.setResponseMessage("OK");
            isOK = true;
        }
        catch (Exception ex) {
            if (isDbgEnabled) {
                LOG.debug("", ex);
            }
            res.setResponseCode("500");
            res.setResponseMessage(ex.toString());
            isOK = false;
        }
        res.sampleEnd();
        res.setSuccessful(isOK);
        return res;
    }
}
