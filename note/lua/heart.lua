function show_love()
    local i, j, k, l, m
    local text = ''
    -- 开头空出5行
     for i=1,5 do
        text = text .. "\n"
    end
    -- 前3行中间有空隙分开来写
    for i = 1, 3, 1 do
        -- 左边的空格，每下一行左边的空格比上一行少2个8*n-2*i
        for j = 1, 32-2*i do
            text = text.." "
        end

        -- 输出左半部分字符小爱心
        for k = 1, 4 * i + 1 do
            text = text.."*"
        end

        -- 中间的空格，每下一行的空格比上一行少4个
        for l=1, 13-4*i do
            text = text.." "
        end

        -- 输出右半部分字符小爱心
        for m=1,4*i+1 do
            text = text.."*"
        end

        -- 每一行输出完毕换行
        text = text.."\n"
    end

    -- 下3行中间没有空格
    for i = 1, 3 do
        -- 左边的空格 8*(n-1)+1
        for j=1,24+1 do
            text = text.." "
        end

        -- 输出字符小爱心
        for k = 1, 29 do
            text = text.."*"
        end

        -- 每一行输出完毕换行
        text = text.."\n"
    end

    -- 下7行
    for i=7, 1, -1 do

        -- 左边的空格，每下一行左边的空格比上一行少2个 8*(n+1)-2*i
        for j=1,40-2*i do
            text = text.." "
        end

        -- 每下一行的字符小爱心比上一行少4个
        for k=1, 4*i-1 do
            text = text.."*"
        end

        -- 每一行输出完毕换行
        text = text.."\n"
    end

    -- 最后一行左边的空格
    for i=1,39 do
        text = text.." "
    end
    -- 最后一个字符小爱心
    text = text.."*\n"

    -- 最后空出5行
    for i=1,5 do
        text = text.."\n"
    end

    return text
end

print(show_love())