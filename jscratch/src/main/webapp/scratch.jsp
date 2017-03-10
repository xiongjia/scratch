<%@ page pageEncoding="UTF-8" %>
<%@ taglib prefix="core" uri="http://java.sun.com/jsp/jstl/core" %>

<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Scratch</title>
</head>
<body>
  <jsp:useBean id="data"  class="scratch.ServletBean">
	<jsp:setProperty name="data" property="testData" value="The test value" />
  </jsp:useBean>

  <h1>Scratch</h1>
  <p>testData = <jsp:getProperty property="testData" name="data"/>
</body>
</html>
