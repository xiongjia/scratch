package xiongjia.sampler;

import java.util.Locale;
import java.util.logging.Logger;

import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNotNull;

import org.apache.jmeter.samplers.SampleResult;
import org.apache.jmeter.util.JMeterUtils;

import xiongjia.sampler.ExampleSampler;

public class TestExampleSampler {
    private final static Logger LOG = Logger.getLogger(TestExampleSampler.class.getName());
    private ExampleSampler exampleSampler;

    @Before
    public void setUp() {
        LOG.info("TestExampleSampler::setUp()");
        JMeterUtils.setLocale(new Locale("ignoreResources"));

        LOG.info("Creating ExampleSampler");
        exampleSampler = new ExampleSampler();
        assertNotNull(exampleSampler);
        exampleSampler.setProperty(ExampleSampler.REP_DATA, "REPData");
    }

    @Test
    public void sample() {
        LOG.info("TestExampleSampler::sample()");
        final SampleResult res = exampleSampler.sample(null);
        assertNotNull(res);
        
        final String sampleData = res.getSamplerData();
        assertEquals(sampleData, "Example Data");
        final String repData = res.getResponseDataAsString();
        assertEquals(repData, "REPData");
        final String repMsg = res.getResponseMessage();
        assertEquals(repMsg, "OK");
    }
}
