# Binding 验证参数

## 1. required
- **描述**：字段是必填的，不能为空。
- **示例**：`Name string `json:"name" binding:"required"`

## 2. len=N
- **描述**：字段的长度必须等于 N。
- **示例**：`Password string `json:"password" binding:"required,len=8"`

## 3. min=N
- **描述**：字段的最小长度为 N。
- **示例**：`Username string `json:"username" binding:"required,min=3"`

## 4. max=N
- **描述**：字段的最大长度为 N。
- **示例**：`Description string `json:"description" binding:"max=100"`

## 5. email
- **描述**：字段必须是有效的电子邮件格式。
- **示例**：`Email string `json:"email" binding:"required,email"`

## 6. url
- **描述**：字段必须是有效的 URL 格式。
- **示例**：`Website string `json:"website" binding:"url"`

## 7. alpha
- **描述**：字段仅允许字母（a-z 和 A-Z）。
- **示例**：`FirstName string `json:"firstName" binding:"alpha"`

## 8. alphanum
- **描述**：字段仅允许字母和数字（a-z、A-Z 和 0-9）。
- **示例**：`Username string `json:"username" binding:"alphanum"`

## 9. printascii
- **描述**：字段仅允许可打印的 ASCII 字符。
- **示例**：`Message string `json:"message" binding:"printascii"`

## 10. oneof=val1,val2,...
- **描述**：字段的值必须在指定的一组值中。
- **示例**：`Gender string `json:"gender" binding:"oneof=male female"`

## 11. gte=N
- **描述**：字段的值必须大于或等于 N。
- **示例**：`Age int `json:"age" binding:"gte=0"` // 年龄需大于或等于 0

## 12. gt=N
- **描述**：字段的值必须大于 N。
- **示例**：`Balance float64 `json:"balance" binding:"gt=0"` // 余额需大于 0

## 13. lte=N
- **描述**：字段的值必须小于或等于 N。
- **示例**：`Rating int `json:"rating" binding:"lte=10"` // 评分需小于或等于 10

## 14. lt=N
- **描述**：字段的值必须小于 N。
- **示例**：`Height float64 `json:"height" binding:"lt=3"` // 身高需小于 3 米

## 15. isbn
- **描述**：字段必须是有效的 ISBN（国际标准书号）。
- **示例**：`ISBN string `json:"isbn" binding:"isbn"`

## 16. uuid
- **描述**：字段必须是有效的 UUID。
- **示例**：`ID string `json:"id" binding:"uuid"`

## 17. creditcard
- **描述**：字段必须是有效的信用卡号。
- **示例**：`CardNumber string `json:"card_number" binding:"creditcard"`
