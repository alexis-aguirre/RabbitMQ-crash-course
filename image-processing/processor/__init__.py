from typing import Dict, Set, Tuple
from uuid import uuid4
from time import sleep

plate_numbers: Set[str] = {
    'ZAG-28-98', 'UKB-043-N', 'NAA-77-41', 'RUG-74-50', '451-TKT-4', 'THM3973', 'FYZ-77-75', 'YWA-24-47', '555-FT',
    'A-32-253', 'YWT6454', 'A22-ALF', 'XVB-92-52', 'AHA-50-15', '48-10-ZP', 'AHN-27-68', 'WZU-45-17', 'FKP-64-97',
    'RLZ9289', 'WZS-57-06', 'MX-10-696'
}

GitHubStorage = "https://github.com/alexis-aguirre/RabbitMQ-crash-course/tree/skeleton/data/license-plates/"


def process_image(data: Dict[str, str]) -> Tuple[bool, Dict[str, str]]:
    image_url: str = data["image"]
    plate_number = image_url.removeprefix(GitHubStorage).removesuffix(".jpg")

    # Legacy app - image processing in the request
    if plate_number not in plate_numbers:
        return False, {}
    sleep(1)
    # Processing stuff ...

    return True, {
        "id": str(uuid4()),
        "location": str(uuid4()),
        "plate_number": plate_number,
    }


__all__ = (
    "process_image",
)
