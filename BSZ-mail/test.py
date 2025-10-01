from os import environ
from exchangelib import Credentials, Account, Configuration, DELEGATE

creds = Credentials(
    username=environ.get("EMAIL_USER"),
    password=environ.get("EMAIL_PASS")
)

config = Configuration(
    server=environ.get("EMAIL_SERVER"),      # your Exchange server
    credentials=creds
)

account = Account(
    primary_smtp_address=environ.get("EMAIL_USER"),
    credentials=creds,
    autodiscover=False,   # since your server is not O365
    config=config,
    access_type=DELEGATE
)

# Test connection
print("Connected! Mailbox:", account.primary_smtp_address)

# Fetch last 5 mails
for item in account.inbox.all().order_by('-datetime_received')[:5]:
    print(item.subject, item.datetime_received)