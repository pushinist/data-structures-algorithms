import enum


class Point:
    def __init__(self, x: float, y: float) -> None:
        self.x: float = x
        self.y: float = y


class orientation_type(enum.Enum):
    collinear: int = 0
    clockwise: int = 1
    counterclockwise: int = 2


def Left_index(points: list[Point]) -> float:
    minn: float = 0
    for i in range(1, len(points)):
        if points[i].x < points[minn].x:
            minn = i
        elif points[i].x == points[minn].x:
            if points[i].y > points[minn].y:
                minn = i
    return minn


def orientation(p: Point, q: Point, r: Point) -> int:
    val: float = (q.y - p.y) * (r.x - q.x) - (q.x - p.x) * (r.y - q.y)

    if val == 0:
        return orientation_type.collinear.value
    elif val > 0:
        return orientation_type.clockwise.value
    else:
        return orientation_type.counterclockwise.value


def convexHull(points: list[Point], n: int) -> None:

    if n < 3:
        return

    left: float = Left_index(points)

    hull: list[float] = []

    p: float = left
    q: float = 0
    while True:

        hull.append(p)

        q: float = (p + 1) % n

        for i in range(n):
            if orientation(points[p], points[i], points[q]) == 2:
                q = i

        p = q

        if p == left:
            break

    for each in hull:
        print(points[each].x, points[each].y)


points: list[Point] = []
points.append(Point(0, 3))
points.append(Point(2, 2))
points.append(Point(1, 1))
points.append(Point(2, 1))
points.append(Point(3, 0))
points.append(Point(0, 0))
points.append(Point(3, 3))

convexHull(points, len(points))
