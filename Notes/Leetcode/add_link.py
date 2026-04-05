import re
import json
import urllib.request
from pathlib import Path

LEETCODE_API = "https://leetcode.com/api/problems/all/"
PATTERN = re.compile(r"\[(\d+)\](?!\()")  # match [123] but not already linked

def fetch_leetcode_map():
    req = urllib.request.Request(
        LEETCODE_API,
        headers={"User-Agent": "Mozilla/5.0"}
    )

    with urllib.request.urlopen(req, timeout=20) as resp:
        data = json.load(resp)

    id_to_slug = {}
    for item in data["stat_status_pairs"]:
        qid = item["stat"]["frontend_question_id"]
        slug = item["stat"]["question__title_slug"]
        id_to_slug[int(qid)] = slug

    return id_to_slug


def format_md_file(file_path: Path, id_to_slug: dict[int, str]):
    content = file_path.read_text(encoding="utf-8")

    def repl(match):
        qid = int(match.group(1))
        slug = id_to_slug.get(qid)

        if not slug:
            return match.group(0)

        name = slug.replace("-", " ")
        url = f"https://leetcode.com/problems/{slug}/"

        return f"[{qid}. {name}]({url})"

    new_content = PATTERN.sub(repl, content)

    if new_content != content:
        file_path.write_text(new_content, encoding="utf-8")
        print(f"Updated: {file_path}")
    else:
        print(f"No change: {file_path}")


def main(folder="."):
    id_to_slug = fetch_leetcode_map()

    for md_file in Path(folder).rglob("*.md"):
        format_md_file(md_file, id_to_slug)


if __name__ == "__main__":
    main(".")


# cd Noted\Leetcode
# python add_link.py