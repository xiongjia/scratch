package wserver.model;

public class MessageStore {
    private final String message;

    public MessageStore() {
        message = "Test Message: test";
    }

    public String getMessage() {
        return message;
    }
}
