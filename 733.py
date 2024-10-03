
from typing import List

def floodFill(image: List[List[int]], sr: int, sc: int, color: int) -> List[List[int]]:
    def fill(img, sr, sc, color, cur):
        if sr < 0 or sr >= len(img) or sc < 0 or sc >=len(img[0]):
            return
        
        if cur != img[sr][sc]: 
            return

        img[sr][sc] = color

        fill(img, sr-1, sc, color, cur)
        fill(img, sr+1, sc, color, cur)
        fill(img, sr, sc+1, color, cur)
        fill(img, sr, sc-1, color, cur)


    if image[sr][sc] == color:
        return image
    
    fill(image, sr, sc, color, image[sr][sc])
    return image

