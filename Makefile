mailer:
	docker build -t mailerservice .
	docker run -p 3002:3002 mailerservice