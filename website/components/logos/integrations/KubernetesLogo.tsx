import { INTEGRATION_WIDTH } from "./constants";

const KubernetesLogo = ({ width = INTEGRATION_WIDTH, className }: { width?: number, className?: string }) => (
    <svg
        viewBox="0 0 32 32"
        xmlns="http://www.w3.org/2000/svg"
        xmlSpace="preserve"
        style={{
            fillRule: "evenodd",
            clipRule: "evenodd",
            strokeLinejoin: "round",
            strokeMiterlimit: 2,
        }}
        className={className || "dark:text-white text-gray-900"}
        width={width}
    >
        <path
            d="m13.604 19.136.011.009-1.333 3.219a6.89 6.89 0 0 1-2.765-3.463l3.437-.584.005.005a.59.59 0 0 1 .645.813v.001Zm-1.109-2.839a.588.588 0 0 0 .229-1.011l.005-.016-2.615-2.339a6.868 6.868 0 0 0-.975 4.339l3.349-.964.007-.009Zm1.526-2.641c.38.276.911.016.932-.453l.016-.005.197-3.495a6.851 6.851 0 0 0-4.016 1.928l2.865 2.025h.006Zm1.015 3.667.964.464.964-.464.239-1.036-.667-.833h-1.072l-.667.833.239 1.036Zm2-4.13a.584.584 0 0 0 .933.447l.009.005 2.844-2.015a6.88 6.88 0 0 0-3.989-1.923l.197 3.485.006.001Zm14.5 7.963-7.697 9.573c-.407.5-1.016.792-1.661.792l-12.349.005c-.645 0-1.26-.292-1.667-.797L.465 21.156a2.118 2.118 0 0 1-.412-1.787L2.804 7.432A2.085 2.085 0 0 1 3.955 6L15.075.683a2.166 2.166 0 0 1 1.848 0l11.125 5.312a2.127 2.127 0 0 1 1.151 1.432L31.95 19.37a2.128 2.128 0 0 1-.412 1.787l-.002-.001Zm-4.385-2.744c-.057-.011-.135-.037-.192-.048-.235-.041-.423-.031-.641-.047-.463-.052-.848-.088-1.192-.197-.141-.052-.24-.219-.287-.292l-.271-.079a8.452 8.452 0 0 0-.141-3.109 8.594 8.594 0 0 0-1.244-2.88c.068-.063.197-.176.233-.213.011-.12 0-.244.125-.375.265-.251.595-.453.989-.699.193-.109.365-.181.557-.323.043-.031.1-.083.147-.12.317-.249.391-.692.161-.979-.23-.287-.672-.312-.989-.063-.047.037-.109.084-.152.12-.176.156-.285.307-.437.469-.328.333-.604.609-.9.807-.125.079-.319.052-.401.047l-.256.183a8.697 8.697 0 0 0-5.525-2.672l-.016-.297c-.088-.083-.192-.156-.219-.333-.031-.359.021-.744.079-1.208.025-.219.077-.396.088-.635v-.188c0-.407-.303-.74-.667-.74-.369 0-.667.333-.667.74v.188c.011.239.063.416.088.635.057.464.105.849.079 1.208a.767.767 0 0 1-.219.344l-.016.281a8.544 8.544 0 0 0-5.552 2.672c-.083-.057-.161-.115-.24-.172-.119.016-.239.052-.395-.036-.297-.204-.573-.48-.901-.813-.151-.161-.26-.312-.437-.463-.043-.037-.104-.084-.147-.12a.844.844 0 0 0-.463-.177.646.646 0 0 0-.532.235c-.229.292-.156.729.161.984l.011.005.141.109c.187.141.359.213.552.323.396.251.724.453.989.699.099.109.12.301.131.385l.213.187a8.642 8.642 0 0 0-1.36 6.011l-.276.079c-.073.099-.177.244-.287.292-.344.109-.729.145-1.192.192-.219.021-.407.011-.641.052-.052.011-.12.032-.177.041l-.004.005h-.011c-.391.095-.647.459-.563.813.077.353.463.572.859.484h.011l.011-.005.172-.036c.229-.063.396-.152.599-.229.437-.156.808-.292 1.161-.344.147-.011.308.093.38.136l.292-.048a8.65 8.65 0 0 0 3.839 4.792l-.12.292c.047.115.095.265.057.375-.125.339-.349.693-.599 1.084-.125.181-.251.323-.36.531-.025.052-.057.131-.083.183-.172.364-.047.787.281.948.333.156.744-.011.921-.38.027-.052.063-.12.084-.172.093-.213.125-.401.192-.609.172-.443.271-.907.516-1.199.068-.077.172-.109.287-.135l.151-.276a8.619 8.619 0 0 0 6.145.015l.141.256c.115.036.24.057.339.208.183.307.307.677.459 1.12.067.208.099.396.192.609.021.047.057.12.084.172.176.369.588.536.916.375.333-.156.459-.579.287-.948a7.273 7.273 0 0 1-.088-.177c-.109-.208-.235-.348-.355-.531-.255-.391-.464-.719-.593-1.057-.052-.172.009-.276.052-.391-.027-.031-.079-.193-.109-.271a8.682 8.682 0 0 0 3.839-4.828c.083.015.233.041.285.052.1-.068.188-.152.371-.141.353.052.724.188 1.161.344.203.079.369.167.599.229.047.016.115.027.172.036l.011.005h.011c.396.089.781-.129.859-.484.084-.355-.172-.719-.563-.812v-.001Zm-5.287-5.48-2.599 2.328v.011a.59.59 0 0 0 .229 1.011l.005.011 3.369.968c.073-.744.027-1.5-.145-2.229a6.87 6.87 0 0 0-.86-2.099l.001-.001Zm-5.348 7.104a.591.591 0 0 0-.537-.312.6.6 0 0 0-.5.312l-1.692 3.057c1.437.491 3 .491 4.437 0l-1.693-3.057h-.015Zm2.515-1.724a.589.589 0 0 0-.646.808v.005l1.344 3.249a6.846 6.846 0 0 0 2.776-3.484l-3.469-.588-.005.01Z"
            fill="currentColor"
            style={{
                fillRule: "nonzero",
            }}
        />
    </svg>
)

export default KubernetesLogo