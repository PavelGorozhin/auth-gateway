import logging
import os
from auth_gateway.config import Config
from auth_gateway.gateway import Gateway
from auth_gateway.utils import setup_logging

def main():
    config = Config()
    setup_logging(config.log_level)
    logger = logging.getLogger(__name__)

    try:
        gateway = Gateway(config)
        gateway.start()
    except Exception as e:
        logger.error(f"Error starting gateway: {str(e)}")
        os._exit(1)

if __name__ == "__main__":
    main()