package constants

const SEND_EMAIL_CODE_TEMPLATE = `
	<div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>
				您于 %s 提交的邮箱验证，本次验证码为:<span>%s</span>
				<br/>
				为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。
			</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>
`
