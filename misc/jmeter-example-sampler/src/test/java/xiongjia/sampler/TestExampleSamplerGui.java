package xiongjia.sampler;

import java.util.Locale;
import java.util.logging.Logger;

import org.junit.Test;
import org.junit.Before;

import static org.junit.Assert.*;

import org.apache.jmeter.util.JMeterUtils;
import org.apache.jmeter.testelement.TestElement;

import xiongjia.sampler.ExampleSamplerGui;
import xiongjia.sampler.ExampleSampler;

public class TestExampleSamplerGui {
    private final static Logger LOG = Logger.getLogger(TestExampleSamplerGui.class.getName()); 
    private ExampleSamplerGui exampleGui;

    @Before
    public void setUp() {
        LOG.info("TestExampleSamplerGui::setUp()");
        JMeterUtils.setLocale(new Locale("ignoreResources"));

        LOG.info("Creating ExampleSamplerGui");
        exampleGui = new ExampleSamplerGui();
        assertNotNull(exampleGui);
    }

    @Test
    public void samplerGui() {
        LOG.info("TestExampleSamplerGui::samplerGui()");
        /* default lab, comment */
        final String samplerLab =  exampleGui.getStaticLabel();
        assertEquals(samplerLab, "My Example");
        final String samplerComment =  exampleGui.getComment();
        assertEquals(samplerComment, "My Example Sampler");
    }

    @Test
    public void testElement() {
        LOG.info("TestExampleSamplerGui::testElement()");
        final TestElement element = exampleGui.createTestElement();
        final String repData =  element.getPropertyAsString(ExampleSampler.REP_DATA);
        assertEquals(repData, "Exmaple Response");
    }
}

