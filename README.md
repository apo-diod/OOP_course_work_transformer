# OOP_course_work_transformer

Has 2 API Endpoints:
1) /add_module -> adds new module and hosts it      
  syntax: 
  {
	  "module": "script",
	  "settings": {
		  "script": "script starting with [def main(data):] that returns dictionary object with transformed data"
	  }
  }
2) /link -> links receiver and transformer modules      
  syntax:
  {
	  "first": "receiver_id",
	  "second": "transformer_id"
  }
