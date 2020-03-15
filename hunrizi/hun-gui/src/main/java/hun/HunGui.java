package hun;

import javafx.application.Application;
import javafx.scene.Scene;
import javafx.scene.control.Alert;
import javafx.scene.control.Button;
import javafx.scene.layout.StackPane;
import javafx.stage.Stage;

public class HunGui extends Application {
  public static void main(String[] args) {
    launch(args);
  }

  @Override
  public void start(Stage primaryStage) throws Exception {
    primaryStage.setTitle("Hello");

    final Button button = new Button();
    button.setText("Test");
    button.setOnAction(event -> {
      final Alert alert = new Alert(Alert.AlertType.INFORMATION);
      alert.setTitle("Info");
      alert.setHeaderText("Header");
      alert.setContentText("Content");
      alert.showAndWait();
    });

    final StackPane rootPane = new StackPane();
    rootPane.getChildren().addAll(button);
    primaryStage.setScene(new Scene(rootPane, 300, 300));
    primaryStage.show();
  }
}
