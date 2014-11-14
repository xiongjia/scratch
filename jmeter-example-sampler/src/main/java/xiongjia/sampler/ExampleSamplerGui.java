package xiongjia.sampler;

import java.awt.BorderLayout;
import java.awt.Component;
import javax.swing.JLabel;
import javax.swing.JPanel;
import javax.swing.JTextArea;

import org.apache.jmeter.samplers.gui.AbstractSamplerGui;
import org.apache.jmeter.testelement.TestElement;
import org.apache.jorphan.logging.LoggingManager;
import org.apache.log.Logger;

import xiongjia.sampler.ExampleSampler;

public class ExampleSamplerGui extends AbstractSamplerGui {
    private static final long serialVersionUID = 3604399572900039865L;

    private static final Logger LOG = LoggingManager.getLoggerForClass();
    private static final boolean isDbgEnabled = LOG.isDebugEnabled();
    
    private static final String samplerLab = "My Example";
    private static final String samplerComment = "My Example Sampler";
    private JTextArea samplerRepData; 

    public ExampleSamplerGui() {
        super();
        if (isDbgEnabled) {
            LOG.debug("ExampleSamplerGui()");
        }

        /* update the sampler panel */
        setComment(samplerComment);
        setLayout(new BorderLayout(0, 5));
        setBorder(makeBorder());
        add(makeTitlePanel(), BorderLayout.NORTH);        
        add(createConfigPanel(), BorderLayout.CENTER);
    }

    private Component createConfigPanel() {
        samplerRepData = new JTextArea();
        samplerRepData.setName(ExampleSampler.REP_DATA);
        samplerRepData.setText("Exmaple Response");

        final JLabel label = new JLabel("Response data:");
        label.setLabelFor(samplerRepData);

        final JPanel dataPanel = new JPanel(new BorderLayout(5, 0));
        dataPanel.add(label, BorderLayout.NORTH);
        dataPanel.add(samplerRepData, BorderLayout.CENTER);
        return dataPanel;
    }

    /**
     * {@inheritDoc}
     */
    @Override
    public String getStaticLabel() {
        return samplerLab;
    }

    /**
     * {@inheritDoc}
     */
    @Override
    public TestElement createTestElement() {
        ExampleSampler sampler = new ExampleSampler();
        this.modifyTestElement(sampler);
        return sampler;
    }

    /**
     * {@inheritDoc}
     */
    @Override
    public String getLabelResource() {
        return this.getClass().getSimpleName();
    }

    /**
     * {@inheritDoc}
     */
    @Override
    public void modifyTestElement(TestElement testElement) {
        testElement.clear();
        this.configureTestElement(testElement);
        testElement.setProperty(ExampleSampler.REP_DATA, samplerRepData.getText());
    }

    /**
     * {@inheritDoc}
     */
    @Override
    public void configure(TestElement element) {
        samplerRepData.setText(element.getPropertyAsString(ExampleSampler.REP_DATA));
        super.configure(element);
    }

    /**
     * {@inheritDoc}
     */
    @Override
    public void clearGui() {
        super.clearGui();
        samplerRepData.setText("Exmaple Response");
    }
}
