import subprocess


def get_incremented_version(version: str) -> str:
    # 버전 문자열이 'v'로 시작하는지 확인
    if not version.startswith('v'):
        raise ValueError("Version string must start with 'v'")

    # 'v'를 제외한 나머지 부분을 숫자 부분으로 분리
    version_numbers = version[1:].split('.')
    # 마지막 숫자를 +1 증가
    version_numbers[-1] = str(int(version_numbers[-1]) + 1)
    # 새로운 버전 문자열을 생성
    return 'v' + '.'.join(version_numbers)


def increment_version(is_remove: bool):
    # tagname이 순으로 내림차순 정렬
    tags = subprocess.run(["git", "tag"], check=True, capture_output=True, text=True).stdout.splitlines()
    if len(tags) == 0:
        return

    last = tags[len(tags) - 1]

    if is_remove:
        for i, tag_name in enumerate(tags):
            subprocess.run(["git", "tag", "-d", tag_name], check=True)
            subprocess.run(["git", "push", "origin", f":refs/tags/{tag_name}"], check=True)

    subprocess.run(["git", "tag", get_incremented_version(last)], check=True)
    subprocess.run(["git", "push", "--tag"], check=True)
