package wserver.action;

import wserver.model.MessageStore;
import com.opensymphony.xwork2.ActionSupport;

public class WServerAction extends ActionSupport {
    private static final long serialVersionUID = 1L;

    private MessageStore messageStore;

    public String execute() {
        messageStore = new MessageStore() ;
        return SUCCESS;
    }

    public MessageStore getMessageStore() {
        return messageStore;
    }

}
