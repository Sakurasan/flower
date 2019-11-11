# coding: utf-8
import sys
 
import cv2
import numpy as np
from keras import models
 
import pretreatment
from mlearn_for_image import preprocess_input
 
 
def get_text(img, offset=0):
    text = pretreatment.get_text(img, offset)
    text = cv2.cvtColor(text, cv2.COLOR_BGR2GRAY)
    text = text / 255.0
    h, w = text.shape
    text.shape = (1, h, w, 1)
    return text
 
 
def main(fn):
    # 读取并预处理验证码
    img = cv2.imread(fn)
    text = get_text(img)
    imgs = np.array(list(pretreatment._get_imgs(img)))
    imgs = preprocess_input(imgs)
 
    # 识别文字
    model = models.load_model('model.v2.0.h5')
    label = model.predict(text)
    label = label.argmax()
    texts = ['打字机', '调色板', '跑步机', '毛线', '老虎', '安全帽', '沙包', '盘子', '本子', '药片', '双面胶', '龙舟', '红酒', '拖把', '卷尺',
             '海苔', '红豆', '黑板', '热水袋', '烛台', '钟表', '路灯', '沙拉', '海报', '公交卡', '樱桃', '创可贴', '牌坊', '苍蝇拍', '高压锅',
             '电线', '网球拍', '海鸥', '风铃', '订书机', '冰箱', '话梅', '排风机', '锅铲', '绿豆', '航母', '电子秤', '红枣', '金字塔', '鞭炮',
             '菠萝', '开瓶器', '电饭煲', '仪表盘', '棉棒', '篮球', '狮子', '蚂蚁', '蜡烛', '茶盅', '印章', '茶几', '啤酒', '档案袋', '挂钟', '刺绣',
             '铃铛', '护腕', '手掌印', '锦旗', '文具盒', '辣椒酱', '耳塞', '中国结', '蜥蜴', '剪纸', '漏斗', '锣', '蒸笼', '珊瑚', '雨靴', '薯条',
             '蜜蜂', '日历', '口哨']
    text = texts[label]
    print(text)
    # 获取下一个词
    # 根据第一个词的长度来定位第二个词的位置
    if len(text) == 1:
        offset = 27
    elif len(text) == 2:
        offset = 47
    else:
        offset = 60
    text = get_text(img, offset=offset)
    if text.mean() < 0.95:
        label = model.predict(text)
        label = label.argmax()
        text = texts[label]
        print(text)
 
    # 加载图片分类器
    model = models.load_model('12306.image.model.h5')
    labels = model.predict(imgs)
    labels = labels.argmax(axis=1)
    for pos, label in enumerate(labels):
        print(pos // 4, pos % 4, texts[label])
 
 
if __name__ == '__main__':
    main(sys.argv[1])
# 运行方式 python3  main.py  <img.jpg>