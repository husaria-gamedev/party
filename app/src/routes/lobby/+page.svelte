<script lang="ts">
    import QRCode from 'qrcode';
    import type { QRCodeToDataURLOptions } from 'qrcode';

    var opts: QRCodeToDataURLOptions = {
        errorCorrectionLevel: 'H',
        type: 'image/png',
        margin: 1,
        scale: 10,
        color: {
            dark: "#000000",
            light: "#FFFFFF"
        }
    }
    let qrData: string = '';

    QRCode.toDataURL('text', opts, (err, url) => {
        if (err) throw err
        qrData = url
    })

</script>

<div class="grid place-items-center h-screen">
	<div class="w-96 max-w-screen p-3">
		<div class="card bg-neutral shadow-xl">
			<figure>
				<img
                class="w-full"
				src="{qrData}"
				alt="Let's Party" />
			</figure>
			<div class="card-body">
				<h2 class="card-title">Lobby</h2>
			</div>
		</div>
    </div>
</div>
